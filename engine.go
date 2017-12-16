package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/smallnest/goreq"
	"github.com/tidwall/gjson"
)

type Engine struct {
	commands map[string]*Command
	cmdmap   map[string]*Command
	noP      map[string]*Command
}

func NewEngine() *Engine {
	return &Engine{
		commands: map[string]*Command{},
		cmdmap:   map[string]*Command{},
		noP:      map[string]*Command{},
	}
}

func (e *Engine) Load(paths ...string) {
	for _, p := range paths {
		e.load(p)
	}
}

func (e *Engine) load(path string) {
	log.Println("Engine Load:", path)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	//	log.Println("Engine Load:read:", string(body))
	commands, err := commandsFromJSON(body)
	if err != nil {
		panic(err)
	}
	//	log.Printf("Engine Load Commands:%+v\n", commands)
	for _, c := range commands {
		if _, has := e.commands[c.Name]; has {
			panic(fmt.Errorf("Load %s Error:Command %s exited.", path, c.Name))
		}
		if _, has := e.cmdmap[c.Name]; has {
			panic(fmt.Errorf("Load %s Error:Command %s exited.", path, c.Name))
		}

		e.cmdmap[c.Name] = c
		if c.Require == "" {
			e.commands[c.Name] = c
		} else {
			e.noP[c.Name] = c
		}
		for _, sc := range c.SubCommand {
			if _, has := e.cmdmap[sc.Name]; has {
				panic(fmt.Errorf("Load %s Error:Command %s exited.", path, sc.Name))
			}
			e.cmdmap[sc.Name] = sc
		}
	}

}

func (e *Engine) Start() {
	log.Printf("Engine Check Commands:%+v\n", e.cmdmap)
	lnoP := len(e.noP)

	for lnoP != 0 {
		for _, c := range e.noP {
			//log.Printf("noP:%+v\n", c)
			if r, ok := e.cmdmap[c.Require]; ok {
				r.SubCommand = append(r.SubCommand, c)
				delete(e.noP, c.Name)
				//	log.Printf("noP Delete:%+v\n", c)
			}
		}

		if len(e.noP) == lnoP {
			panic(fmt.Errorf("Commands [%+v] don't find Require.\n", e.noP))
		}
		lnoP = len(e.noP)
	}

	for k := range e.cmdmap {
		delete(e.cmdmap, k)
	}

	log.Printf("Engine Start Commands:%+v\n", e.commands)
	for _, c := range e.commands {
		context := NewContext()
		e.Exec(nil, context, c)
	}
}

func (e *Engine) Exec(req *goreq.GoReq, context *Context, cmd *Command) {
	log.Printf("Engine Exec:%+v\n", cmd)
	if req == nil {
		req = goreq.New()
		req.Debug = true
	}

	switch cmd.Method {
	case "POST", "post", "p", "P":
		req.Post(context.P(cmd.URL))
	case "DELETE", "delete", "d", "D":
		req.Delete(context.P(cmd.URL))
	default:
		req.Get(context.P(cmd.URL))
	}

	for k, v := range cmd.Header {
		req.SetHeader(context.P(k), context.P(v))
	}

	if cmd.ContentType != "" {
		req.ContentType(cmd.ContentType)
	}

	paramstr := context.P(string(*cmd.Params))
	req.Query(paramstr)

	_, body, errs := req.EndBytes()
	if len(errs) != 0 {
		panic(errs[0])
	}

	// resp := map[string]interface{}{}
	// err := json.Unmarshal(body, &resp)
	// if err != nil {
	// 	log.Println("Engine Exec Resp Error:", string(body), err)
	// 	panic(err)
	// }

	gjsons := gjson.ParseBytes(body)

	//log.Printf("Engine Exec Resp:%+v\n", resp)

	//TODO 只能上下文确认返回值。。。
	for k, v := range cmd.Return {
		kp := context.P(k)
		if vp, ok := v.(string); ok {
			v = context.P(vp)
		}
		rv := gjsons.Get(kp)
		if !rv.Exists() {
			panic(fmt.Errorf("Resp key %s[%s] don't exists.", k, kp))
		}
		rvi := rv.Value()
		if reflect.TypeOf(v).Name() == reflect.TypeOf(rvi).Name() && fmt.Sprint(v) == fmt.Sprint(rvi) {
			continue
		}
		log.Println("checkvalue:", reflect.TypeOf(v).Name(), reflect.TypeOf(rvi).Name(), v, rvi)
		panic(fmt.Errorf("Key:%s[%s] Value:%v != %v", k, kp, v, rvi))
	}
	for k, v := range cmd.Context {
		kp := context.P(k)
		v = context.P(v)
		rv := gjsons.Get(kp)
		if !rv.Exists() {
			panic(fmt.Errorf("Resp key %s[%s] don't exists.", k, kp))
		}
		context.K(kp, rv.Value())
	}

	for _, c := range cmd.SubCommand {
		e.Exec(req, context, c)
	}
}
