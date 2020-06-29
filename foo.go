package main

import "fmt"

// Foo
type Foo struct {
	Message string
}

// NewFoo
func NewFoo(message string) *Foo {
	return &Foo{Message: message}
}

// Greet
func (f *Foo) Greet() string {
	return fmt.Sprintf("Hello %s", f.Message)
}
