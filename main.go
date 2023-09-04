package main

import (
	"fmt"
	"strconv"
	"strings"
	"learn-go/cube"
	"learn-go/circle"
)

func main() {
	fmt.Println("======= String class =======")

	var name string = "messi"
	formattedPi := fmt.Sprintf("Value: %.2f", 3.14)
	fmt.Printf("Hello %s \n", name)
	fmt.Println("Pi " + formattedPi)

	const myName string = "rifki"
	const piValue float32 = 3.14
	fmt.Printf("Hello " + myName + "\n")
	fmt.Printf("Pi value %f \n", piValue)

	const (
		square          = "kotak"
		isToday bool    = true
		numeric uint8   = 1
		floatNum        = 2.2
	)

	fmt.Printf("Bentuk adalah " + square + "\n")
	fmt.Printf("hari adalah %s \n", strconv.FormatBool(isToday))

	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("number %d", i)
		fmt.Printf("Print %s \n", str)
	}

	var i int = 0

	for i < 5 {
		fmt.Printf("Formatted print %d \n", i)
		i++
	}

	var message = []string{"Thank", "you =======\n"}
    printMessage("======= String class finished,", message)

	/**
	 * Doing some math
	 */

	fmt.Println("======= Math class =======")

	var area, circumference = circle.AreaAndCircumference(5.0)
	fmt.Printf("Circle area %.2f, Circle circumference %.2f \n", area, circumference)

	var rectangleArea = cube.Area(1, 2, 3)
	fmt.Printf("Cube area %.2f \n", rectangleArea)

	var cubeVolume = cube.Volume(2.00, cube.Area)
	fmt.Printf("Cube volume %.2f \n", cubeVolume)

	fmt.Println("======= Math class finished, Thank you =======")

}

func printMessage(message string, messages []string) {
	var joinedMessage = getString(messages)
	fmt.Println(message, joinedMessage)
}

func getString(messages []string) string {
	return strings.Join(messages, " ")
}
