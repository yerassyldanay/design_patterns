package main

import "fmt"

type ContextPassenger interface {
	Next()
}

type Context struct {
	i int
	handlers []HandlerFunc
}

func (c *Context) Next() {
	if c.i >= len(c.handlers) {
		return
	}

	err := c.handlers[c.i](c)
	if err != nil {
		return
	}

	c.i++
	c.Next()
}

type HandlerFunc func(*Context) error

func PointAFunc(ctx *Context) error {
	fmt.Println("passed through point A")
	return nil
}

func PointBFunc(ctx *Context) error {
	fmt.Println("passed through point B")
	return nil
}

func PointSuccessFunc(ctx *Context) error {
	fmt.Println("passed through point Success")
	return nil
}

func main()  {
	var c = &Context{
		i:        0,
		handlers: []HandlerFunc{
			PointAFunc,
			PointBFunc,
			PointSuccessFunc,
		},
	}
	c.Next()
}

