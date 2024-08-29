package main

import "github.com/longtime1116/jitsuyo_go/u"

// main() よりも先に呼ばれる。ここで設定を登録してmainで最初に実行する・・・とかすることができる
func init() {
	u.P("init() called")
}

func main() {
	u.P("main() called")
}
