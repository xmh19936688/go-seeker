package main

import (
	"fmt"

	"github.com/xmh19936688/go-seeker/seeker"
)

func main() {
	str := `hour
minute
second
day
week
month
year`

	lines := seeker.New(str)
	if lines.SeekToString("day") {
		fmt.Println("processing:", lines.CurrentLine())
	}
}
