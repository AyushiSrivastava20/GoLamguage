package main

import "fmt"

// practice go basics
type Books struct {
	tittle string
	author string
	price  float64
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)

}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	fmt.Println("***Struct***")
	var Book1 Books

	Book1.tittle = "go program"
	Book1.author = "me"
	Book1.price = 23.54
	fmt.Printf("book details: %s\n", Book1.tittle)
	fmt.Println("***factorial***")
	var i int = 10
	fmt.Printf("factorial of %d is %d", i, factorial(i))

	fmt.Println("***Factorial end***")
	fmt.Println("***fibonacci***")
	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d ", fibonaci(j))
	}

	fmt.Println("*********")
	nameThing := map[string]string{"Ayushi": "neu", "xyz": "rit", "mark": "mit"}
	for name := range nameThing {
		fmt.Println("name", name, "school", nameThing[name])
	}
	name, ok := nameThing["Ayushi"]
	if ok {
		fmt.Println("found", name)
	} else {
		fmt.Println("not found")
	}

}
