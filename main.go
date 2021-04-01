package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func printTimeOver() {
	timeOver := `
.___________. __  .___  ___.  _______      ______   ____    ____  _______ .______      
|           ||  | |   \/   | |   ____|    /  __  \  \   \  /   / |   ____||   _  \     
.---|  |----.|  | |  \  /  | |  |__      |  |  |  |  \   \/   /  |  |__   |  |_)  |    
    |  |     |  | |  |\/|  | |   __|     |  |  |  |   \      /   |   __|  |      /     
    |  |     |  | |  |  |  | |  |____    |   --   |    \    /    |  |____ |  |\  \----.
    |__|     |__| |__|  |__| |_______|    \______/      \__/     |_______|| _|  ._____|`

	fmt.Println(timeOver)
}

func printTimeOver2() {
	timeOver := `
___ _ _  _ ____    ____ _  _ ____ ____ 
 |  | |\/| |___    |  | |  | |___ |__/ 
 |  | |  | |___    |__|  \/  |___ |  \ `

	fmt.Println(timeOver)
}

func printTimeOver3() {
	timeOver := `
████████╗██╗███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
╚══██╔══╝██║████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
   ██║   ██║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
   ██║   ██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
   ██║   ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
   ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝ `
	fmt.Println(timeOver)
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	// put the digits (placeholders) into variables
	// using the placeholder array type above
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	//fmt.Println(len(os.Args))

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <countdown in minutes>")
		fmt.Println("Usage: countdown.exe <countdown in minutes>")
		return
	}

	mins, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("You must enter a the countdown in minutes as an integer!")
	}

	//fmt.Println(mins)

	// convert the input to a time array
	endtime := time.Now().Local().Add(time.Minute * time.Duration(mins))

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		// get the difference to the end time of the countdown
		diff := getTimeRemaining(endtime)

		if diff.t <= 0 {
			screen.Clear()
			//fmt.Println("Time Over")
			printTimeOver3()
			return
		}

		//hour, min, sec := now.Hour(), now.Minute(), now.Second()
		hour, min, sec := diff.h, diff.m, diff.s

		// extract the digits: 17 becomes, 1 and 7 respectively
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			// Print a line for each placeholder in clock
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}

			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
