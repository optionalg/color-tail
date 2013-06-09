package technicolor

import (
	"math/rand"
)

var (
	colorTable = map[string]string{
		"black": "\033[30m",
		"red": "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"magenta": "\033[35m",
		"cyan": "\033[36m",
		"white": "\033[37m",
		"reset": "\033[39m",
	}

	Colors = []string{ "black" , "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	ctLen = int32(len(Colors))
)

// Colorizes given string with specified color
// if the color doesn't exist we fallback to a random one
func Paint(str, color string) string {

	c, ok := colorTable[color]
	if !ok {
		return RandPaint(str)
	} else {
		return c + str + colorTable["reset"]
	}
}

// Colorizes the string with randomly picked color
func RandPaint(str string) string {
	idx := rand.Int31n(ctLen)
	key := Colors[idx]
	return Paint(str, key)
}