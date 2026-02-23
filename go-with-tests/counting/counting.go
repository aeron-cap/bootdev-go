package counting

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleep struct{
	Calls int
}

func (d *DefaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper DefaultSleep) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, "Go!")
}

func main() {
	sleeper := &DefaultSleep{}
	Countdown(os.Stdout, *sleeper)
}