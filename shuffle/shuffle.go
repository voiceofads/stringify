package shuffle

import (
	"math/rand"
	"time"
)

func Shuffle(s string) string {

	rand.Seed(time.Now().Unix())

	inRune := []rune(s)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
