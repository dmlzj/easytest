package main

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
