package TextAnimator

import (
	"fmt"
	"time"
)

type TextAnimator struct {
	preloaderParts string
	textTime       time.Duration
}

func (t TextAnimator) PrintSequential(text string) {

	var str string
	var preloadCounter int

	for i, char := range text {
		str += string(char)

		if i+1 != len(text) && text[i+1] == 10 {
			fmt.Printf("\r%s", str)
			str = "\r"
		} else if i+1 == len(text) {
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

func NewTextAnimator(textTime time.Duration) TextAnimator {
	preloadParts := "-\\|/"

	t := TextAnimator{
		preloaderParts: preloadParts,
		textTime:       textTime,
	}

	return t
}
