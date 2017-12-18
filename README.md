# easytest
API接口自动化测试框架
## 配置文件
`Header`内数据用于所有的用例请求，但优先级低于用例内`header`,`Value`为全局上下文。
```
{
    "Header":{
        "X-Appkey":"weixin"
    },
    "Value":{
        "HOST":"http://192.168.1.9:18080"
    }
}
```
## 用例文件
测试用例可写在多个文件内，程序会自动构建`Require`树。
```go
type Command struct {
	Name        string                 `json:"name" description:"名称"`
	URL         string                 `json:"url" description:"请求地址"`
	Method      string                 `json:"method" description:"请求方式"`
	Require     string                 `json:"require" description:"前置需求"`
	ContentType string                 `json:"contenttype" description:"ContentType"`
	RequestLua  []string               `json:"requestlua" description:"请求前调用的lua文件"`
	Header      map[string]string      `json:"header" description:"请求头"`
	Params      *json.RawMessage       `json:"params" description:"请求参数"`
	Return      map[string]interface{} `json:"return" description:"期望返回"`
	NextLua     []string               `json:"nextjs" description:"执行后续命令前调用的lua文件"`
	Context     map[string]string      `json:"context" description:"上下文"`
	SubCommand  []*Command             `json:"subcommand" description:"子命令"`
}
```
```json
{
    "name": "shops",
    "require":"login",
    "url": "http://192.168.1.200:9000/shops",
    "method": "get",
    "header": {
        "X-Appkey": "weixin"
    },
    "params": {
        "fields": "shop_name"
    },
    "return": {
        "code": 0,
        "data.nickname": "nzlov"
    }
}
```
## 用例语法
### Path Syntax
路径是由一个点分隔的一系列键。一个键可能包含特殊的通配符`*`和`?`。要访问数组值，请使用索引作为键。要获取数组中的元素数量或访问子路径，请使用`#`字符。点和通配符可以用`\`来转义。
```
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}
```
```
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Deer Hunter"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig"
```
您还可以使用`#[…`或找到与`#[…]#`的所有匹配。查询支持`==`,`!=`,`<`,`<=`,`=>`,`>`比较运算符和简单模式匹配`%`运算符。
```
friends.#[last=="Murphy"].first    >> "Dale"
friends.#[last=="Murphy"]#.first   >> ["Dale","Jane"]
friends.#[age>45]#.last            >> ["Craig","Murphy"]
friends.#[first%"D*"].last         >> "Murphy"
```
### 引入上下文语法
在配置文件中使用`{{.name}}`形式标记需要引用的上下文语法，程序会自动替换。
### 请求前后事件处理
内置`lua`脚本,按照配置文件中配置的先后顺序依次调用。
* 前置脚本传入数据有`Context`,`Cmd`,`Req`
* 后置脚本传入数据有`Context`,`Cmd`