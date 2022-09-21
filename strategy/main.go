package main

import "fmt"

type Strategy interface {
	Execute()
}

type ConcreteStrategy1 struct{}

func (c *ConcreteStrategy1) Execute() {
	fmt.Println("ConcreteStrategy1 - execute()")
}

type ConcreteStrategy2 struct{}

func (c *ConcreteStrategy2) Execute() {
	fmt.Println("ConcreteStrategy2 - execute()")
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SomeMethodThatUsesStrategy() {
	c.strategy.Execute()
}

func (c *Context) ChangeStrategy(strategy Strategy) {
	c.strategy = strategy
}

func main() {
	c := NewContext(&ConcreteStrategy1{})
	c.SomeMethodThatUsesStrategy()

	c.ChangeStrategy(&ConcreteStrategy2{})
	c.SomeMethodThatUsesStrategy()
}
