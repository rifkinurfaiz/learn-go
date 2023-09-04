package structs

type person struct {
	Name string
	Age int
}

type Employee struct {
	Grade int
	person
}
