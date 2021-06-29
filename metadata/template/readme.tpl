# 淘宝开放平台 golang SDK


## 编译工具链
```sh
make tools
```
- 在bin目录下会生成downloader和generator两个可执行文件 
- 因为使用embed特性，编译downloder和generator需要golang版本>= 1.16


## 下载淘宝开放平台ApiMetadata
```sh
./bin/downloader -meta=./metadata/assets/json
```

### downloader 参数
- meta: 下载metadata目录
- pkg: 指定下载特定包，如：-pkg=user, 仅下载"用户API"

## 根据metadata生成API Golang SDK
```sh
./bin/generator -meta=./metadata/assets/json -patch=./metadata/assets/patch
```

### generator 参数
- meta: metadata文件所在目录
- patch: metadata patch文件所在目录
- pkg: 指定生成特定API包 


### API 调用
```golang
package main

import (
    "log"

    "github.com/bububa/opentaobao/core"
    userModel "github.com/bububa/opentaobao/model/user"
    userApi "github.com/bububa/opentaobao/api/user"
)

func main() {
    clt := core.NewSDKClient(APP_KEY, APP_SECRET)
    req := userModel.NewTaobaoUserAvatarGetRequest()
    req.SetNick("nick")
    resp, err := userApi.TaobaoUserAvatarGet(clt, req)
    if err != nil {
        log.Fatalln(err)
    }
    log.Printf("%+v\n", resp)
}
```

## 常见问题

### API分包
为避免生成最终目录内文件过多，以及命名空间冲突问题，根据淘宝开放平台API分类原则对API分包。
分包设置在metadata/assets/package.json，"pkg"参数即为API包名，与"id"(API树节点ID)对应。
下载的metadata如果在package.json无法找到对应的包则不会下载保存。

### 字段类型映射
| 淘宝字段类型 | go SDK 类型 |
| ----------- | -----------|
| Number | int64 |
| Price | float64 |
| Boolean | bool |
| String | string |
| Date | string |
| Json | string |
| Field List | []string |
| 其他 | struct指针 |

- 如果字段类型后带"[]", 则转换为slice
- 如果字段UsePointer=true, 则转换为指针
- 同一保重struct名称一样会对struct成员进行合并

淘宝API中对象命名空间非常混乱，即使在同一包中仍然会出现struct冲突的情况。如果出现struct命名空间冲突有两种解决办法:
- 修改metadata/assets/conflict_models.json文件，增加可能冲突的struct名。generator会对该struct前增加API名进行重命名，比如 
AlibabaWdkFinanceOrderBackflow API中的ApiResult会被重命名为AlibabaWdkFinanceOrderBackflowApiResult
- 也可增加对应的patch文件修改字段的Type。generator在生成API时会先查找对应的API有没有patch，如果有patch则使用patch文件替换metadata文件

### 缺少特定API
对于无法通过downloader下载的API metaddata, 可以参照已下载metadata文件在patch目录自行添加相应的metadata, generator会生成相应的API。

## API 列表
| 淘宝API分类 | 对应SDK package |
| ---------- | ---------- |
{{- range $v := .Pkgs }}
| [x] [{{ $v.Name }}]({{ $v.Link }}) | [github.com/bububa/opentaobao/api/{{ $v.Pkg }}](https://pkg.go.dev/github.com/bububa/opentaobao/api/{{ $v.Pkg }}) |
{{- end  }}

