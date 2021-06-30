package lc

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	res := []string{}
	sub := generateParenthesis(n - 1)
	for _, v := range sub {
		res = append(res, fmt.Sprintf("(%s)", v), "()"+v)
		left := "()" + v
		right := v + "()"
		if left != right {
			res = append(res, right)
		}
	}
	return res
}

func Test() {
	fmt.Println(generateParenthesis(3))

}
