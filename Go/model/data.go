package model

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	fmt.Println("hola mundo")
}

//s := Person{Name: "Sean", Age: 50}
//s := Person{}
//s.Name = "Sean"
//s.Age = 42
