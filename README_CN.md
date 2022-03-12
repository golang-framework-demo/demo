# DEMO - golang framework

## 目录结构
```
|-- projects(demo)
|   |-- src
|   |   |-- app
|   |   |   |-- db
|   |   |   |   |-- models
|   |   |   |   |   `-- ... ...
|   |   |   |   |-- repositories
|   |   |   |   |   `-- ... ...
|   |   |   |-- middleware
|   |   |   |   `-- middleware.go
|   |   |   |-- services
|   |   |   |   |-- demo
|   |   |   |   |   |-- controller
|   |   |   |   |   |   `-- (controller)demo_controller.go
|   |   |   |   |   |-- service
|   |   |   |   |   |   `-- (service)demo_service.go
|   |   |   |   |-- ... ...
|   |   |-- autoload
|   |   |   `-- autoload.go
|   |   |   `-- autoload_before.go
|   |   |-- storage
|   |   |   `-- error.go
|   `-- .dev.yaml
|   `-- .router.yaml
|   `-- main.go
|   `-- go.mod
`-- LICENCE
`-- README.md
`-- README_CN.md
```

## 快速入门

#### 进入 ./demo/src 文件根目录
```sh
$ cd ./demo/src
```

#### 启动配置 yaml 文件, .(dev|stg|prd).yaml
```yaml
Common:
  Name: "golang mvc framework demo" # Project Name {项目名称}
  Port: "8577"                      # Port {启动端口}
  TimeLocation: "Asia/Shanghai"     # Time Zone {时区}
  Addr: ""                          # URL {项目url地址}
  Hssl:                             # HTTPS 配置
    Power: 0                        # 0 - 不启用 | 1 - 启用
    CertFile: ""                    # 证书地址
    KeysFile: ""                    # 密钥地址

db:                                 # 数据库配置
  requirement:
    - Driver: "mysql"               # 数据库引擎
      MaxOpen: 50                   # 最大链接数
      MaxIdle: 100                  # 最大空闲数
      MaxLifeTime: 60               # 最大生命周期
      ShowedSQL: true               # 是否在控制器中显示sql语句
      CachedSQL: false              # 是否开启缓存
      Expire: 10                    # 缓存开启后、缓存时间
      MaxElementSize: 1000          # 最大可保存元素大小
      TimeLoaction: "Asia/Shanghai" # 数据库时区
  conns:
    - Require: 0                    # 用reqirement中那个基础数据引擎配置
      Repo: "admin_platform_go"     # 库名
      Host: "127.0.0.1"             # 链接地址
      Port: 3306                    # 链接端口
      Username: "root"              # 用户名
      Password: "root"              # 密码

redis:
  - Network: "tcp"                  # 连接方式
    Addr: "127.0.0.1:6379"          # 连接地址
    Username: "root"                # 用户名
    Password: ""                    # 密码
    DB: 0                           # Redis库
    PoolSize: 20                    # Redis缓存池
```

###1. 开发调试
```sh
$ go mod tidy
$ go mod vendor
$ go run -mod=vendor main.go --env=dev #(dev|stg|prd) 对应 .(dev|stg|prd).yaml 文件
```

###2. 测试&生产发布
#### 生成二进制文件
```sh
$ go mod tidy
$ go mod vendor
$ go build -mod=vendor -o demo main.go
```
#### 启动二进制文件
```sh
$ ./demo --env=dev 
```
备注：.*.yaml 文件需要放在与二进制文件相同的根目录下

###3. 启动成功
```sh
[golang-mvc framework] v1.0.0
[golang-mvc framework] Listening and serving HTTP on :8577
```

## 路由配置

#### 路由 yaml 文件配置 .router.yaml
```yaml
route:
  Tag:
    Demo: &demo golang_framework_demo # 服务模块标签
  Rel:
    *demo: "s_demo" # 对应服务模块对应路由组
```

#### 路由启动代码
```sh
$ cat ./demo/src/autoload/autoload.go
```
```go
package autoload

import (
	"github.com/golang-framework/mvc"
	err "github.com/golang-framework/mvc/modules/error"
	"src/storage"
)

type autoload struct {
}

func init() { (&autoload{}).src(mvc.New()) // 启动 golang-framework 框架  }

func (ad *autoload) src(fw *mvc.Framework) {
	ad.before() // 框架启动前, 需启动的功能

	// initialized self error message
	fw.Err = &err.M{EMsg: storage.Em} // 自定义错误信息载入
	fw.Fw()

	// initialized router
	fw.Route.E, fw.Route.M = ad.mvcInitializedRouter() // 路由配置载入
	fw.FwRouter()

	fw.Run() // 启动
}
```

```sh
$ cat ./demo/src/autoload/autoload_before.go
```
```go
package autoload

import (
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/routes"
	"src/app/middleware"
	"src/app/services/demo/controller"
)

func (ad *autoload) before() {

}

func (ad *autoload) mvcInitializedRouter() (*routes.AHC, *routes.M) {
	var (
		demo = controller.NewDemoController() // 初始化DEMO控制器
	)

	return &routes.AHC{middleware.NoRouter(), middleware.NoMethod()}, // 自定义无路由返回
		&routes.M{
		    // 对应.router.yaml配置服务模块标签
			property.Instance.Get("route.Tag.Demo", "").(string): {
				// 路由前置过滤中间件
				Middleware: &routes.AHC{},
				// 路由适配器
				Adapter: map[*routes.I]*routes.AHC{
					// 路由配置:{ 对应路由控制器 }
					// routes.Ai("{_}") - 默认路由 - {router_group[s_demo]}/{controller_name[demo]}/{action_name[test_demo]} - 默认方式 GET
					// routes.Ai("{_}", http.MethodPost) - 默认路由 - 自定义方式 POST
					// routes.Ai("_self_define_t", http.MethodGet) - 自定义路由 - 自定义方式
					// - 自定义路由地址 - {router_group[s_demo]}/_self_define_t
					// 下属路由地址: http://127.0.0.1:8577/s_demo/demo/test_demo
					routes.Ai("{_}"): {demo.TestDemo},
				},
			},
		}
}
```

```sh
$ cat ./demo/src/main.go
```
```go
package main

import (
	_ "src/autoload" // 自动载入
)
func main() {}
```
