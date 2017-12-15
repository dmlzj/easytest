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
### return 取值
* 层级用`.`连接
* 数组用`:`标记。eg: data.list:0.shopid //表示`data`下的`list`数组中的第一个元素的`shopid`
### 引入上下文语法
在配置文件中使用`{{name}}`形式标记需要引用的上下文语法，程序会自动替换。
### 请求前后事件处理
内置`js`脚本,按照配置文件中配置的先后顺序依次调用。
* 前置脚本传入数据有`context`,`req`
* 后置脚本传入数据有`context`