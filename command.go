package main

import (
	"encoding/json"
)

type Command struct {
	Name        string                 `json:"name" description:"名称"`
	URL         string                 `json:"url" description:"请求地址"`
	Method      string                 `json:"method" description:"请求方式"`
	Require     string                 `json:"require" description:"前置需求"`
	ContentType string                 `json:"contenttype" description:"ContentType"`
	RequestJS   []string               `json:"requestjs" description:"请求前调用的js文件"`
	Header      map[string]string      `json:"header" description:"请求头"`
	Params      *json.RawMessage       `json:"params" description:"请求参数"`
	Return      map[string]interface{} `json:"return" description:"期望返回"`
	NextJS      []string               `json:"nextjs" description:"执行后续命令前调用的js文件"`
	Context     []string               `json:"context" description:"上下文"`
	SubCommand  []*Command             `json:"subcommand" description:"子命令"`
}
