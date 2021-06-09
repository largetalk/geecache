package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/etcd-io/etcd/clientv3"
	"wzq.com/geecache"
)

const geecache_etcd_prefix = "/geecache/nodes"

type GeeEtcClient struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

func Newclient(config clientv3.Config) (*GeeEtcClient, error) {
	c, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	kv := clientv3.NewKV(c)
	lease := clientv3.NewLease(c)
	return &GeeEtcClient{
		client: c,
		kv:     kv,
		lease:  lease,
	}, nil
}

func (c *GeeEtcClient) RegisteNode(addr string) (bool, error) {
	key := fmt.Sprintf(
		"%v/%v",
		geecache_etcd_prefix,
		addr,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	_, err := c.client.Put(ctx, key, addr)
	defer cancel()
	if err != nil {
		log.Println("put error", err)
		return false, err
	}

	_, cancel = context.WithTimeout(context.Background(), time.Minute)
	resp, err := c.client.Get(context.TODO(), key)
	defer cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return false, err
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return true, nil
}

func (c *GeeEtcClient) Watch(peers *geecache.HTTPPool) {
	rch := c.client.Watch(context.Background(), geecache_etcd_prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			peers.Add(string(ev.Kv.Value))
		}
	}
}
