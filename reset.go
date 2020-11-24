package fortune_telling

import "time"

var (
	resetTime  = time.Date(0, 0, 0, 23, 30, 0, 0, time.Local)
	resetTimer = time.NewTicker(1 * time.Minute)
)

func Reset() {
	signedMap = make(map[string]int, 0)
}

func SetTime(reset time.Time) {
	resetTime = reset
}

func doReset() {
	for {
		select {
		case curr := <-resetTimer.C:
			if curr.Hour() == resetTime.Hour() && curr.Minute() == resetTime.Minute() {
				Reset()
			}
		}
	}
}
