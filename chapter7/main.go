package main

func main() {

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
