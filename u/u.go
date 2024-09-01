package u

import (
	"fmt"

	secret "github.com/longtime1116/jitsuyo_go/u/internal"
)

func P(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
func PP(a any) (n int, err error) {
	return fmt.Printf("%+v\n", a)
}

func Pf(format string, a ...any) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

// internal の確認用(chapter6)
func PublicSecretP(a ...any) (n int, err error) {
	return secret.SecretP(a...)
}

// 変数未利用によるコンパイルエラー回避用
func N(a ...any) {
}
