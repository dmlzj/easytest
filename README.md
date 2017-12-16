# easytest
API接口自动化测试框架
## 配置文件
测试用例可写在多个文件内，程序会自动构建`Require`树。
```go
type Command struct {
	Name        string                 `json:"name" description:"名称"`
	URL         string                 `json:"url" description:"请求地址"`
	Method      string                 `json:"method" description:"请求方式"`
	Require     string                 `json:"require" description:"前置需求"`
	ContentType string                 `json:"contenttype" description:"ContentType"`
	RequestJS   []string               `json:"requestjs" description:"请求前调用的js文件"`
	Header      map[string]string      `json:"header" description:"请求头"`
	Params      map[string]interface{} `json:"params" description:"请求参数"`
	Return      map[string]interface{} `json:"return" description:"期望返回"`
	NextJS      []string               `json:"nextjs" description:"执行后续命令前调用的js文件"`
	Context     []string               `json:"context" description:"返回值保存到上下文"`
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
## 配置语法
### Path Syntax
A path is a series of keys separated by a dot. A key may contain special wildcard characters '*' and '?'. To access an array value use the index as the key. To get the number of elements in an array or to access a child path, use the '#' character. The dot and wildcard characters can be escaped with '\'.
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
You can also query an array for the first match by using #[...], or find all matches with #[...]#. Queries support the ==, !=, <, <=, >, >= comparison operators and the simple pattern matching % operator.
```
friends.#[last=="Murphy"].first    >> "Dale"
friends.#[last=="Murphy"]#.first   >> ["Dale","Jane"]
friends.#[age>45]#.last            >> ["Craig","Murphy"]
friends.#[first%"D*"].last         >> "Murphy"
```
### 引入上下文语法
在配置文件中使用`{{.name}}`形式标记需要引用的上下文语法，程序会自动替换。
### 请求前后事件处理
内置`js`脚本,按照配置文件中配置的先后顺序依次调用。
* 前置脚本传入数据有`context`,`req`
* 后置脚本传入数据有`context`