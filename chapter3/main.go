package main

import (
	"fmt"

	"github.com/longtime1116/jitsuyo_go/u"
)

type Person struct {
	Name string
	Age  int
}
type Student struct {
	Person
	School string
}

func (p Person) Hello() {
	fmt.Printf("Hello, I'm %v.\n", p.Name)
}
func (s Student) Hello2() {
	// s.Name のように、Personを経由しなくても呼び出せる
	fmt.Printf("Hello, I'm %v. I'm a student at %v.\n", s.Name, s.School)
}

func String[T any](s T) string {
	return fmt.Sprintf("%v", s)
}

type Struct struct {
	t interface{}
}

func (s Struct) String() string {
	return fmt.Sprintf("%v", s.t)
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

	s := Struct{"hoge"}
	u.P(String(s))
	u.P(s.String())

	s1 := Student{
		Person: Person{Name: "Tom", Age: 20},
		School: "MIT",
	}
	fmt.Println(s1)
	s1.Hello() // Person のメソッドも呼び出せる
	s1.Hello2()
}
