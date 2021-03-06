module main

go 1.16

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/etcd-io/etcd v3.3.25+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	wzq.com/geecache v0.0.0
	wzq.com/lc v0.0.0
)

replace wzq.com/geecache => ./geecache
replace wzq.com/lc => ./lc

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
