package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) SayHello() {
    fmt.Println("Hello, I am", p.Name)
}

func (p *Person) Grow() {
    p.Age++
}

func main() {
    scores := map[string]int{
        "Bob": 90,
        "Alice": 85,
        "Serhii": 100,
		"Stels": 99,
    }
    scores["Tom"] = 70
    fmt.Println(scores)

	user := Person{Name: "Serhii", Age: 34}
    user.SayHello()
    user.Grow()
    user.Grow()
    fmt.Println("Age now:", user.Age)

}