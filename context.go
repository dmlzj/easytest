package main

import (
	"bytes"
	"html/template"
)

type Context struct {
	value map[string]interface{}
}

func NewContext() *Context {
	return &Context{
		value: map[string]interface{}{},
	}
}

func (c *Context) K(k string, v interface{}) {
	c.value[k] = v
}
func (c *Context) V(k string) (interface{}, bool) {
	v, ok := c.value[k]
	return v, ok
}
func (c *Context) P(str string) string {
	//log.Printf("Context P:%s:%+v\n", str, c.value)
	t, err := template.New("context").Parse(str)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBufferString("")
	err = t.Execute(body, c.value)
	if err != nil {
		panic(err)
	}
	return body.String()
}
