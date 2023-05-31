package TextAnimator

import (
	"fmt"
	"time"
)

type TextAnimator struct {
	text           string
	spinnerParts   []rune
	textChannel    chan string
	spinnerChannel chan rune
}

func (t TextAnimator) textAnimate() {
	defer close(t.textChannel)
	var str string

	for _, char := range t.text {
		str = str + string(char)
		t.textChannel <- str
		fmt.Printf("\r%s", str)
		time.Sleep(100 * time.Millisecond)
	}
}

func (t TextAnimator) spinner(delay time.Duration) {
	defer close(t.spinnerChannel)
	for {
		for _, r := range t.spinnerParts {
			t.spinnerChannel <- r
			time.Sleep(delay)
		}
	}
}

func NewTextAnimator(text string, spinnerParts []rune) TextAnimator {
	t := TextAnimator{
		text:           text,
		spinnerParts:   spinnerParts,
		textChannel:    make(chan string),
		spinnerChannel: make(chan rune),
	}

	return t
}

func (t TextAnimator) PrintSequential() {

	var str string

	go t.textAnimate()
	go t.spinner(100 * time.Millisecond)

	for str = range t.textChannel {
		fmt.Printf("\r%s%c", str, <-t.spinnerChannel)
	}

	_, ok := <-t.textChannel

	if !ok {
		fmt.Printf("\r%s%s", str, " \n")
	}

	time.Sleep(5 * time.Second)
}
