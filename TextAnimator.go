package TextAnimator

import (
	"fmt"
	"time"
)

type TextAnimator struct {
	text           string
	preloaderParts string
	textTime       time.Duration
}

func (t TextAnimator) PrintSequential() {

	var str string
	var preloadCounter int

	for _, char := range t.text {
		str += string(char)
		fmt.Printf("\r%s%c", str, t.preloaderParts[preloadCounter])
		if char == 10 {
			str = "\r"
		}

		if preloadCounter >= len(t.preloaderParts)-1 {
			preloadCounter = 0
		} else {
			preloadCounter++
		}
		time.Sleep(t.textTime)
	}
}

func NewTextAnimator(text string) TextAnimator {
	preloadParts := "-\\|/"

	t := TextAnimator{
		text:           text,
		preloaderParts: preloadParts,
		textTime:       75 * time.Millisecond,
	}

	return t
}
