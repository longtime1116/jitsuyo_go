package u

import (
	"fmt"
)

func P(a ...any) (n int, err error) {
	return fmt.Println(a)
}

func Pf(format string, a ...any) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

// 変数未利用によるコンパイルエラー回避用
func N(a ...any) {
}
