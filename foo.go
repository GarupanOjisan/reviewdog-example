package main

import "fmt"

type Foo struct {
	Message string
}

func NewFoo(message string) *Foo {
	return &Foo{Message: message}
}

func (f *Foo) Greet() string {
	return fmt.Sprintf("Hello %s", f.Message)
}
