package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/smallnest/goreq"
)

type Engine struct {
	commands map[string]*Command
}

func NewEngine() *Engine {
	return &Engine{
		commands: map[string]*Command{},
	}
}

func (e *Engine) Load(path string) error {
	log.Println("Engine Load:", path)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	//	log.Println("Engine Load:read:", string(body))
	commands, err := commandsFromJSON(body)
	if err != nil {
		return err
	}
	//	log.Printf("Engine Load Commands:%+v\n", commands)
	for _, c := range commands {
		if _, has := e.commands[c.Name]; has {
			return fmt.Errorf("Command %s exited.", c.Name)
		}
		e.commands[c.Name] = &c
	}
	return nil
}

func (e *Engine) Start() []error {
	//	log.Printf("Engine Start Commands:%+v\n", e.commands)
	for _, c := range e.commands {
		context := NewContext()
		errs := e.Exec(nil, context, c)
		if len(errs) != 0 {
			return errs
		}
	}
	return []error{}
}

func (e *Engine) Exec(req *goreq.GoReq, context *Context, cmd *Command) []error {
	log.Printf("Engine Exec:%+v\n", cmd)
	if req == nil {
		req = goreq.New()
		req.Debug = true
	}

	switch cmd.Method {
	case "GET", "get", "g", "G":
		req.Get(cmd.URL)
	case "POST", "post", "p", "P":
		req.Post(cmd.URL)
	case "DELETE", "delete", "d", "D":
		req.Delete(cmd.URL)
	}

	for k, v := range cmd.Header {
		req.SetHeader(k, v)
	}

	if cmd.ContentType != "" {
		req.ContentType(cmd.ContentType)
	}

	params := map[string]interface{}{}
	for k, v := range cmd.Params {
		switch t := v.(type) {
		case string:
			params[k] = t
		default:
			params[k] = v
		}
	}
	pbody, err := json.Marshal(params)
	if err != nil {
		return []error{err}
	}
	req.Query(string(pbody))

	_, body, errs := req.EndBytes()
	if len(errs) != 0 {
		return errs
	}

	resp := map[string]interface{}{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []error{err}
	}

	log.Printf("Engine Exec Resp:%+v\n", resp)

	contexts := map[string]struct{}{}
	for _, v := range cmd.Context {
		contexts[v] = struct{}{}
	}

	for k, v := range cmd.Return {
		rv, err := checkvalue(strings.Split(k, "."), v, resp)
		if err != nil {
			return []error{err}
		}
		if _, ok := contexts[k]; ok {
			context.K(k, rv)
		}
	}

	for _, c := range cmd.SubCommand {
		err := e.Exec(req, context, c)
		if err != nil {
			return err
		}
	}
	return nil
}
