package u

import (
	"fmt"

	secret "github.com/longtime1116/jitsuyo_go/u/internal"
)

func P(a ...any) (n int, err error) {
	return fmt.Println(a...)
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

// cmt+shift+p をして、 go interface stubs とかやって出てきたやつを使うと、stub関数を作ってくれる
// ここでは、MyInterfaceとMyStructを自分で作っていて、MyInterfaceを満たすべき関数群を自動で作ってくれている
type MyInterface interface {
	DoSomething(a int) error
	DoSomethingElse(b string) string
}

func (x MyStruct) DoSomething(a int) error {
	panic("not implemented") // TODO: Implement
}

func (x MyStruct) DoSomethingElse(b string) string {
	panic("not implemented") // TODO: Implement
}

type MyStruct struct{}
