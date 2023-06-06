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

	for i, char := range t.text {
		str += string(char)

		if i+1 != len(t.text) && t.text[i+1] == 10 {
			fmt.Printf("\r%s", str)
			str = "\r"
		} else if i+1 == len(t.text) {
			fmt.Printf("\r%s", str)
		} else {
			fmt.Printf("\r%s%c", str, t.preloaderParts[preloadCounter])
		}

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
