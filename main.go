package main

import "fmt"

type User struct {
	A string
	B int
}

func main()  {
	u := User{
		A: "",
		B: 0,
	}
	fmt.Println("hello reviewdog")
	fmt.Println(u.A)
}
