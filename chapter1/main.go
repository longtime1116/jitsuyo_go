package main

import (
	"github.com/longtime1116/jitsuyo_go/u"
)

// go における定数は、コンパイル時に値が定まりバイナリに埋め込まれる。
const (
	a = 1
	b = 1 + 2          // 計算可能
	c = 1 << 100       // 定数は明示的に型をしていない限りは型に束縛されないので、このような大きい数値も定義可能
	d = "Hello, World" // もちろん文字列なども
)

type BirthCountry uint64

const (
	Japan BirthCountry = iota + 1
	China
	US
)

type Language uint64

const (
	C Language = 1 << iota
	Java
	Golang
	Ruby
	Python
)

type Person struct {
	country  BirthCountry
	language Language
}

const ()

func main() {
	u.P("chapter1")
	u.Pf("%v", 100)
	u.P(a, b, c>>99)

	Me := Person{
		country:  Japan,
		language: C | Golang | Ruby,
	}
	u.P(Me)
	u.Pf("%b", Me.language)

	// 可能な限り短縮系を使ったほうがいいが、以下はどっちでも良い
	var hoge1 int8 = 3
	hoge2 := int8(3)
	u.N(hoge1, hoge2)

}
