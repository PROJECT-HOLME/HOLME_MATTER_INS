package features

import (
	"INS_LIGHTBULB/src/terminal"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - state: bool -> shows whether the light bulb is on or off
  - degree: int -> shows the degree (percentage) of light power.
  - color: string -> shows the color of the light

- Return Value: string
	- It will return a completely formatted message.
*/

func returnFormattedMsg(state bool, degree int, color string) string {
	if !state {
		return "Light state: OFF"
	} else {
		return "Light state: ON\nCurrent Light Degree: " + strconv.Itoa(degree) + "%\nCurrent Light Color: " + color
	}
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - none

- Return Value: []byte
	- It returns byte array which is printable.
*/
func readImageFile(isOn bool, figNum int) []byte {
	if !isOn || figNum%2 == 0 {
		buf, err := os.ReadFile("src/raw/off.txt")
		if err != nil {
			panic(err)
		}
		return buf
	} else {
		buf, err := os.ReadFile("src/raw/on.txt")
		if err != nil {
			panic(err)
		}
		return buf
	}
}

// PrintLightDisabled
/*
- It prints default light.
- It will read a image of light bulb, and print in default state.
- Parameter:
  - none

- Return Value: void
	- It will do its task and exit the function.
*/
func PrintLightDisabled() {
	buf := readImageFile(false, 0)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.HiBlack(string(buf))
	color.HiBlack(returnFormattedMsg(false, 0, ""))
}

func PrintLightEnabled(degree int, flag uint8) {
	var j int
	if degree > 0 {
		j = 4
	} else {
		j = 3
	}
	for i := 1; i < j; i++ {
		buf := readImageFile(true, i)
		terminal.ClearTerminal() //* Clear the terminal first.

		//* Print will be differed by the flag.
		//* Flag will be 8bit unsigned integer.
		switch flag {
		case 0:
			{ //* Default Light
				color.White(string(buf))
				color.White(returnFormattedMsg(true, degree, "WHITE"))
			}
		case 1:
			{ //* Yellow Light
				color.Yellow(string(buf))
				color.Yellow(returnFormattedMsg(true, degree, "YELLOW"))
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(1000 * time.Millisecond)
}
