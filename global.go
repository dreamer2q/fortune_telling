package fortune_telling

import (
	"encoding/json"
	"math/rand"
	"time"
)

const (
	fileJson = "signs.json"
)

var (
	signs     []Telling
	signedMap map[string]int
	levelMap  = map[string]int{
		"下下": 0,
		"中下": 1,
		"中平": 2,
		"中吉": 3,
		"上吉": 4,
		"上上": 5,
		"大吉": 6,
	}
)

func init() {
	err := json.Unmarshal([]byte(signData), &signs)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())
	Reset()
	go doReset()
}
