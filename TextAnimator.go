package TextAnimator

import (
	"fmt"
	"time"
)

type TextAnimator struct {
	text             string
	preloaderParts   string
	textChannel      chan string
	preloaderChannel chan rune
	preloaderTime    time.Duration
	textTime         time.Duration
}

func (t TextAnimator) textAnimate() {
	defer close(t.textChannel)
	var str string

	for _, char := range t.text {
		str = str + string(char)
		t.textChannel <- str
		fmt.Printf("\r%s", str)
		time.Sleep(t.textTime)
	}
}

func (t TextAnimator) preloader() {
	defer close(t.preloaderChannel)
	for {
		for _, r := range t.preloaderParts {
			t.preloaderChannel <- r
			time.Sleep(t.preloaderTime)
		}
	}
}

func (t TextAnimator) PrintSequential() {

	var str string

	go t.textAnimate()
	go t.preloader()

	for str = range t.textChannel {
		fmt.Printf("\r%s%c", str, <-t.preloaderChannel)
	}

	_, ok := <-t.textChannel

	if !ok {
		fmt.Printf("\r%s%s", str, " \n")
	}

	time.Sleep(5 * time.Second)
}

func NewTextAnimator(text string) TextAnimator {
	preloadParts := "-\\|/"

	t := TextAnimator{
		text:             text,
		preloaderParts:   preloadParts,
		textChannel:      make(chan string),
		preloaderChannel: make(chan rune),
		preloaderTime:    100 * time.Microsecond,
		textTime:         50 * time.Millisecond,
	}

	return t
}
