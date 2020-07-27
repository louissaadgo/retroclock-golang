package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for {
		hour := time.Now().Hour()
		min := time.Now().Minute()
		sec := time.Now().Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				// colon blink
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
		screen.MoveTopLeft()
		screen.Clear()
	}
}
