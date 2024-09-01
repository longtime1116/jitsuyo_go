package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/longtime1116/jitsuyo_go/u"
)

type Player struct {
	Name  string    `json:"name"`
	Wins  int       `json:"wins"`
	Birth time.Time `json:"birth"`
}

func main() {
	f, err := os.Open("player.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var player Player
	// io.Reader インタフェースを満たしている型の(ストリーミングなJSONデータを)デコードをするときは、Decoder()が便利。
	if err := json.NewDecoder(f).Decode(&player); err != nil {
		log.Fatal(err)
	}
	u.PP(player)

	// []byte 型を扱う場合にはjson.Unmarshal()
	s := `{ "name": "Bob", "wins": 3, "birth": "1992-12-30T04:34:56+09:00" }`
	if err := json.Unmarshal([]byte(s), &player); err != nil {
		log.Fatal(err)
	}
	u.P(player)

}
