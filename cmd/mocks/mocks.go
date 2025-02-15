package mocks

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleep struct{}

func (d *DefaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

/**
*
* sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
* Countdown(os.Stdout, sleeper)
*
**/
