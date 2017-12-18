package main

import (
	"bytes"
	"html/template"
	"log"
)

type Context struct {
	Value  map[string]interface{}
	Header map[string]string
}

func NewContext() *Context {
	return &Context{
		Value:  map[string]interface{}{},
		Header: map[string]string{},
	}
}
func NewContextWithCopy(c *Context) *Context {
	m := map[string]interface{}{}
	for k, v := range c.Value {
		m[k] = v
	}
	h := map[string]string{}
	for k, v := range c.Header {
		h[k] = v
	}
	return &Context{
		Value:  m,
		Header: h,
	}
}

func (c *Context) K(k string, v interface{}) {
	log.Println("Context K:", k, v)
	c.Value[k] = v
}
func (c *Context) V(k string) (interface{}, bool) {
	v, ok := c.Value[k]
	return v, ok
}
func (c *Context) HK(k, v string) {
	log.Println("Context HK:", k, v)
	c.Header[k] = v
}
func (c *Context) HV(k string) (string, bool) {
	v, ok := c.Header[k]
	return v, ok
}
func (c *Context) P(str string) string {
	//log.Printf("Context P:%s:%+v\n", str, c.value)
	t, err := template.New("context").Parse(str)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBufferString("")
	err = t.Execute(body, c.Value)
	if err != nil {
		panic(err)
	}
	return body.String()
}
