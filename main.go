package main

import (
	"fmt"
	"learn-go/shapes/circle"
	"learn-go/shapes/cube"
	"learn-go/shapes"
	"learn-go/structs"
	"learn-go/shapes/rectangle"
	"strconv"
	"strings"
	"runtime"
)

func main() {
	/**
	 * ==========================================
	 * 			   String manipulation
	 * ==========================================
	 */
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
	 * ========================================
	 * 				Doing some math
	 * ========================================
	 */
	fmt.Println("======= Math class =======")

	var area, circumference = circle.AreaAndCircumference(5.0)
	fmt.Printf("Circle area %.2f, Circle circumference %.2f \n", area, circumference)

	var rectangleArea = cube.Area(1, 2, 3)
	fmt.Printf("Cube area %.2f \n", rectangleArea)

	var cubeVolume = cube.Volume(2.00, cube.Area)
	fmt.Printf("Cube volume %.2f \n", cubeVolume)

	fmt.Print("======= Math class finished ======= \n\n")

	/**
	 * ==========================================
	 * 					OOP
	 * ==========================================
	 */
	fmt.Println("======= OOP class =======")

	type student struct {
		name string
		grade int
	}

	student1 := student{}
	student1.name = "messi"
	student1.grade = 5

	var student2 = student{"anto", 4}

	fmt.Printf("student 1 name: %s \n", student1.name)
	fmt.Printf("student 1 grade: %s \n", strconv.Itoa(student1.grade))

	fmt.Printf("student 2 name: %s \n", student2.name)
	fmt.Printf("student 2 grade: %s \n\n", strconv.Itoa(student1.grade))

	var employee1 = structs.Employee{}
	employee1.Name = "charles"
	employee1.Age = 25
	employee1.Grade = 10

	fmt.Printf("employee 1 name: %s \n", employee1.Name)
	fmt.Printf("employee 1 age: %s \n", strconv.Itoa(employee1.Age))
	fmt.Printf("employee 1 grade: %s \n\n", strconv.Itoa(employee1.Grade))

	var rect = rectangle.Rectangle{Width: 5, Height: 3}
	var circle = circle.Circle{Radius: 2}

	// Use the Shape interface to calculate the area and perimeter of both shapes.
	var shapes = []shapes.Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Printf("shape type: %T\n", shape)
		fmt.Printf("area: %.2f\n", shape.Area())
		fmt.Printf("perimeter: %.2f\n", shape.Perimeter())
	}

	fmt.Print("======= OOP class finished ======= \n\n")

	/**
	 * ==========================================
	 * 					GOROUTINE
	 * ==========================================
	 */
	fmt.Println("======= GOROUTINE class =======")

	runtime.GOMAXPROCS(2)

	go print(5, "Hello there")
	print(5, "Apa kabar")

	var input float32
	fmt.Scanln(&input)

	fmt.Print("======= GOROUTINE class finished =======\n\n")

	fmt.Println("======= Channel class =======")

	runtime.GOMAXPROCS(2)

    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
    }

    go sayHelloTo("charles leclerc")
    go sayHelloTo("carlos sainz")
    go sayHelloTo("lewis hamilton")

    fmt.Println(<-messages)

    var message2 = <-messages
    fmt.Println(message2)

    var message3 = <-messages
    fmt.Println(message3)

	fmt.Print("======= Channel class finished =======\n\n")

}

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func printMessage(message string, messages []string) {
	var joinedMessage = getString(messages)
	fmt.Println(message, joinedMessage)
}

func getString(messages []string) string {
	return strings.Join(messages, " ")
}
