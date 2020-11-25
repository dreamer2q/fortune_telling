package fortune_telling

import (
	"errors"
	"math/rand"
)

//Telling the fortune result struct
type Telling struct {
	Level   string
	Content string
	Detail1 string
	Detail2 string
}

//Ask ask for a sign with key,
//if key has asked, err is not nil
func Ask(key string) (Telling, error) {
	defer func() { signedMap[key]++ }()
	if !HasAsked(key) {
		return genTelling(), nil
	}
	return Telling{}, errors.New("already asked for a fortune-telling")
}

//HasAsked check whether the key has asked
func HasAsked(key string) bool {
	return signedMap[key] > 0
}

func genTelling() Telling {
	var index = rand.Intn(len(signs))
	return signs[index]
}

//String return stars the Level indicates
func (t Telling) String() string {
	return LevelStars(t.Level)
}

//LevelStars return a stars string matched with the level
func LevelStars(l string) string {
	var (
		sep = levelMap[l]
		ret string
	)
	for i := 0; i < 6; i++ {
		var cell string
		if i < sep {
			cell = "★"
		} else {
			cell = "☆"
		}
		ret += cell
	}
	return ret
}
