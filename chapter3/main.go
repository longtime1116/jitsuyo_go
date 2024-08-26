package main

import "github.com/longtime1116/jitsuyo_go/u"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 後々フィールドが追加された時のことを考えると、フィールド名は省略しない方が望ましいという判断をしても良い。
	me := Person{
		Name: "M",
		Age:  121,
	}
	var me2 = Person{"M2", 19}
	u.N(me)
	u.N(me2)

}
