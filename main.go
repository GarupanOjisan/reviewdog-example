package main

import "fmt"

type User struct {
	A string
	B int
}

func main()  {
	foo := NewFoo("foo")
	fmt.Println(foo.Greet())
}
