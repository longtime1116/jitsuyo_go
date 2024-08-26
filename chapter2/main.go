package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type People []Person

func (p People) Adult() People {
	ret := make([]Person, 0, len(p))
	for _, v := range p {
		if v.Age >= 20 {
			ret = append(ret, v)
		}
	}
	return ret
}

func (p Person) String() string {
	return fmt.Sprintf("[String()]   %v(%v) in %v", p.Name, p.Age, "xxxx")
}

func (p Person) GoString() string {
	return fmt.Sprintf("[GoString()] %v(%v) in %v", p.Name, p.Age, "xxxx")
}

func (p Person) Debug() string {
	return fmt.Sprintf("[Debug()]    %v(%v) in %v", p.Name, p.Age, p.Address)
}

func main() {

	// sliceにも型名をつけられる
	p := People{
		{"Abby", 9, "Yokohama"},
		{"Bob", 23, "Tokyo"},
		{"Charly", 44, "Kobe"},
	}

	// String() と GoString() をうまくやることで、マスキングできる。
	for _, v := range p.Adult() {
		fmt.Println(v)
		fmt.Printf("%#v\n", v)
		fmt.Println(v.Debug())
	}
}
