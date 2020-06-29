package main

import "fmt"

// Foo just foo
type Foo struct {
	Message string
}

// NewFoo creating a new foo
func NewFoo(message string) *Foo {
	return &Foo{Message: message}
}

// Greet by foo
func (f *Foo) Greet() string {
	return fmt.Sprintf("Hello %s", f.Message)
}
