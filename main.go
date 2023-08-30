package main

import (
	"fmt"
	"strconv"
	"strings"
	"learn-go/utils"
)

func main() {
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

	for i:= 0; i < 5; i++ {
		str := fmt.Sprintf("number %d", i)
		fmt.Printf("Print %s \n", str)
	}

	var i int = 0

	for i < 5 {
		fmt.Printf("Formatted print %d \n", i)
		i++
	}

	var message = []string{"Thank", "you\n\n"}
    printMessage("Finished,", message)

	var area, circumference = utils.Calculate(5.0)
	fmt.Printf("Circle area %.2f, Circle circumference %.2f \n", area, circumference)

	var rectangleArea = utils.Area(1, 2, 3, 4)
	fmt.Printf("Rectangle area %.2f", rectangleArea)
}

func printMessage(message string, messages []string) {
	var joinedMessage = getString(messages)
	fmt.Println(message, joinedMessage)
}

func getString(messages []string) string {
	return strings.Join(messages, " ")
}