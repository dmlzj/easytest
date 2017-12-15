package main

type Command struct {
	Name        string                 `json:"name" description:"名称"`
	URL         string                 `json:"url" description:"请求地址"`
	Method      string                 `json:"method" description:"请求方式"`
	ContentType string                 `json:"contenttype" description:"ContentType"`
	Header      map[string]string      `json:"header" description:"请求头"`
	Params      map[string]interface{} `json:"params" description:"请求参数"`
	Return      map[string]interface{} `json:"return" description:"期望返回"`
	Context     []string               `json:"context" description:"上下文"`
	SubCommand  []*Command             `json:"subcommand" description:"子命令"`
}
