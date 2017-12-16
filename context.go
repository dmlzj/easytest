package main

import (
	"bytes"
	"html/template"
	"log"
)

type Context struct {
	value  map[string]interface{}
	header map[string]string
}

func NewContext() *Context {
	return &Context{
		value:  map[string]interface{}{},
		header: map[string]string{},
	}
}
func NewContextWithCopy(c *Context) *Context {
	m := map[string]interface{}{}
	for k, v := range c.value {
		m[k] = v
	}
	return &Context{
		value: m,
	}
}

func (c *Context) K(k string, v interface{}) {
	log.Println("Context K:", k, v)
	c.value[k] = v
}
func (c *Context) V(k string) (interface{}, bool) {
	v, ok := c.value[k]
	return v, ok
}
func (c *Context) HK(k, v string) {
	log.Println("Context HK:", k, v)
	c.header[k] = v
}
func (c *Context) HV(k string) (string, bool) {
	v, ok := c.header[k]
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
