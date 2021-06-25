# 淘宝开放平台 golang SDK


## 编译工具链
```sh
make tools
```
* 在bin目录下会生成downloader和generator两个可执行文件 


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


## API 列表

- [x] 阿里通信API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/alicom(https://pkg.go.dev/github.com/bububa/opentaobao/api/alicom)

- [x] 阿里健康医(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/alihealth(https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealth)

- [x] 物联API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/alink(https://pkg.go.dev/github.com/bububa/opentaobao/api/alink)

- [x] 阿里大于API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/aliqin(https://pkg.go.dev/github.com/bububa/opentaobao/api/aliqin)

- [x] 阿里云API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/aliyun(https://pkg.go.dev/github.com/bububa/opentaobao/api/aliyun)

- [x] 阿里云ocsAPI(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/aliyunocs(https://pkg.go.dev/github.com/bububa/opentaobao/api/aliyunocs)

- [x] 本地生活API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/alsc(https://pkg.go.dev/github.com/bububa/opentaobao/api/alsc)

- [x] 反欺诈风控API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/antifraud(https://pkg.go.dev/github.com/bububa/opentaobao/api/antifraud)

- [x] 司法拍卖(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/auction(https://pkg.go.dev/github.com/bububa/opentaobao/api/auction)

- [x] 阿里车联网API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/autonavi(https://pkg.go.dev/github.com/bububa/opentaobao/api/autonavi)

- [x] 百川(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/baichuan(https://pkg.go.dev/github.com/bububa/opentaobao/api/baichuan)

- [x] 宝点API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/baodian(https://pkg.go.dev/github.com/bububa/opentaobao/api/baodian)

- [x] 保险API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/baoxian(https://pkg.go.dev/github.com/bububa/opentaobao/api/baoxian)

- [x] 淘宝内容API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/beehive(https://pkg.go.dev/github.com/bububa/opentaobao/api/beehive)

- [x] 账务API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/bill(https://pkg.go.dev/github.com/bububa/opentaobao/api/bill)

- [x] 会员卡(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/blackvip(https://pkg.go.dev/github.com/bububa/opentaobao/api/blackvip)

- [x] 汽车票API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/bus(https://pkg.go.dev/github.com/bububa/opentaobao/api/bus)

- [x] 彩票API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/caipiao(https://pkg.go.dev/github.com/bububa/opentaobao/api/caipiao)

- [x] 旅行用车API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/car(https://pkg.go.dev/github.com/bububa/opentaobao/api/car)

- [x] 类目API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/category(https://pkg.go.dev/github.com/bububa/opentaobao/api/category)

- [x] 淘宝抽奖平台API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/choujiang(https://pkg.go.dev/github.com/bububa/opentaobao/api/choujiang)

- [x] 淘宝同城API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/cityretail(https://pkg.go.dev/github.com/bububa/opentaobao/api/cityretail)

- [x] 百川推送(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/cloudpush(https://pkg.go.dev/github.com/bububa/opentaobao/api/cloudpush)

- [x] yunos推送服务api(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/cmns(https://pkg.go.dev/github.com/bububa/opentaobao/api/cmns)

- [x] 菜鸟配送API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/cntms(https://pkg.go.dev/github.com/bububa/opentaobao/api/cntms)

- [x] 店铺会员管理API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/crm(https://pkg.go.dev/github.com/bububa/opentaobao/api/crm)

- [x] DMP API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/dmp(https://pkg.go.dev/github.com/bububa/opentaobao/api/dmp)

- [x] 阿里健康药API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/drug(https://pkg.go.dev/github.com/bububa/opentaobao/api/drug)

- [x] 数据API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/dt(https://pkg.go.dev/github.com/bububa/opentaobao/api/dt)

- [x] 生活汇API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/elife(https://pkg.go.dev/github.com/bububa/opentaobao/api/elife)

- [x] 电子物流API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/eticket(https://pkg.go.dev/github.com/bububa/opentaobao/api/eticket)

- [x] 分销API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/fenxiao(https://pkg.go.dev/github.com/bububa/opentaobao/api/fenxiao)

- [x] 电影票API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/film(https://pkg.go.dev/github.com/bububa/opentaobao/api/film)

- [x] 机票API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/flight(https://pkg.go.dev/github.com/bububa/opentaobao/api/flight)

- [x] 淘宝游戏API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/game(https://pkg.go.dev/github.com/bububa/opentaobao/api/game)

- [x] 游戏激励平台API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/gameact(https://pkg.go.dev/github.com/bububa/opentaobao/api/gameact)

- [x] 菜鸟无线API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/guoguo(https://pkg.go.dev/github.com/bububa/opentaobao/api/guoguo)

- [x] 酒店导购API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/hotel(https://pkg.go.dev/github.com/bububa/opentaobao/api/hotel)

- [x] 国际站商品API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/icbu(https://pkg.go.dev/github.com/bububa/opentaobao/api/icbu)

- [x] 天猫互动接口(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/interact(https://pkg.go.dev/github.com/bububa/opentaobao/api/interact)

- [x] 智能设备(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/iot(https://pkg.go.dev/github.com/bububa/opentaobao/api/iot)

- [x] JAE(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/jae(https://pkg.go.dev/github.com/bububa/opentaobao/api/jae)

- [x] ONS消息服务(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/jms(https://pkg.go.dev/github.com/bububa/opentaobao/api/jms)

- [x] 聚石塔API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/jst(https://pkg.go.dev/github.com/bububa/opentaobao/api/jst)

- [x] 聚划算API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/ju(https://pkg.go.dev/github.com/bububa/opentaobao/api/ju)

- [x] 知识库API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/kclub(https://pkg.go.dev/github.com/bububa/opentaobao/api/kclub)

- [x] 地动仪(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/lbs(https://pkg.go.dev/github.com/bububa/opentaobao/api/lbs)

- [x] 法务诉讼对外API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/legalcase(https://pkg.go.dev/github.com/bububa/opentaobao/api/legalcase)

- [x] 生活服务API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/lifeservice(https://pkg.go.dev/github.com/bububa/opentaobao/api/lifeservice)

- [x] 物流API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/logistic(https://pkg.go.dev/github.com/bububa/opentaobao/api/logistic)

- [x] 码上淘API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/ma(https://pkg.go.dev/github.com/bububa/opentaobao/api/ma)

- [x] 多媒体平台API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/media(https://pkg.go.dev/github.com/bububa/opentaobao/api/media)

- [x] 天猫美妆API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/mei(https://pkg.go.dev/github.com/bububa/opentaobao/api/mei)

- [x] 喵街API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/mos(https://pkg.go.dev/github.com/bububa/opentaobao/api/mos)

- [x] 手机淘宝API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/mtop(https://pkg.go.dev/github.com/bububa/opentaobao/api/mtop)

- [x] 手淘开放API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/mtopopen(https://pkg.go.dev/github.com/bububa/opentaobao/api/mtopopen)

- [x] 网上法庭对外API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/nazca(https://pkg.go.dev/github.com/bububa/opentaobao/api/nazca)

- [x] 文本算法API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/nlp(https://pkg.go.dev/github.com/bububa/opentaobao/api/nlp)

- [x] openimAPI(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/openim(https://pkg.go.dev/github.com/bububa/opentaobao/api/openim)

- [x] 拍卖API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/paimai(https://pkg.go.dev/github.com/bububa/opentaobao/api/paimai)

- [x] 商品API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/product(https://pkg.go.dev/github.com/bububa/opentaobao/api/product)

- [x] 营销API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/promotion(https://pkg.go.dev/github.com/bububa/opentaobao/api/promotion)

- [x] 千牛接口(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/qianniu(https://pkg.go.dev/github.com/bububa/opentaobao/api/qianniu)

- [x] 奇门仓储API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/qimen(https://pkg.go.dev/github.com/bububa/opentaobao/api/qimen)

- [x] 质检品控API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/qt(https://pkg.go.dev/github.com/bububa/opentaobao/api/qt)

- [x] 退款API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/refund(https://pkg.go.dev/github.com/bububa/opentaobao/api/refund)

- [x] 国际站外贸直通车API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/scbp(https://pkg.go.dev/github.com/bububa/opentaobao/api/scbp)

- [x] 聚安全API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/security(https://pkg.go.dev/github.com/bububa/opentaobao/api/security)

- [x] 服务平台API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/servicecenter(https://pkg.go.dev/github.com/bububa/opentaobao/api/servicecenter)

- [x] 店铺API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/shop(https://pkg.go.dev/github.com/bububa/opentaobao/api/shop)

- [x] 直通车API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/simba(https://pkg.go.dev/github.com/bububa/opentaobao/api/simba)

- [x] 子账号管理API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/subuser(https://pkg.go.dev/github.com/bububa/opentaobao/api/subuser)

- [x] Tanx API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tanx(https://pkg.go.dev/github.com/bububa/opentaobao/api/tanx)

- [x] 虚拟院线API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/taotv(https://pkg.go.dev/github.com/bububa/opentaobao/api/taotv)

- [x] 淘宝客API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tbk(https://pkg.go.dev/github.com/bububa/opentaobao/api/tbk)

- [x] 门票-商品管理API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/ticket(https://pkg.go.dev/github.com/bububa/opentaobao/api/ticket)

- [x] 天猫会员积分(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmallcms(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallcms)

- [x] 天猫国际API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmallhk(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallhk)

- [x] 天猫精品库API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmallitem(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallitem)

- [x] 天猫服务数据API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmallsc(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallsc)

- [x] 天猫服务商品API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmallservice(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallservice)

- [x] 消息服务API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tmc(https://pkg.go.dev/github.com/bububa/opentaobao/api/tmc)

- [x] 交易API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/trade(https://pkg.go.dev/github.com/bububa/opentaobao/api/trade)

- [x] 评价API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/traderate(https://pkg.go.dev/github.com/bububa/opentaobao/api/traderate)

- [x] 火车票API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/train(https://pkg.go.dev/github.com/bububa/opentaobao/api/train)

- [x] 度假-商品管理API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/travel(https://pkg.go.dev/github.com/bububa/opentaobao/api/travel)

- [x] tv支付(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/tvpay(https://pkg.go.dev/github.com/bububa/opentaobao/api/tvpay)

- [x] 用户API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/user(https://pkg.go.dev/github.com/bububa/opentaobao/api/user)

- [x] 工具API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/util(https://pkg.go.dev/github.com/bububa/opentaobao/api/util)

- [x] 旺旺API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wangwang(https://pkg.go.dev/github.com/bububa/opentaobao/api/wangwang)

- [x] 电子面单API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/waybill(https://pkg.go.dev/github.com/bububa/opentaobao/api/waybill)

- [x] 五道口API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wdk(https://pkg.go.dev/github.com/bububa/opentaobao/api/wdk)

- [x] 手淘分享(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wirelessshare(https://pkg.go.dev/github.com/bububa/opentaobao/api/wirelessshare)

- [x] 物流宝API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wlb(https://pkg.go.dev/github.com/bububa/opentaobao/api/wlb)

- [x] 菜鸟集货API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wlbimports(https://pkg.go.dev/github.com/bububa/opentaobao/api/wlbimports)

- [x] 菜鸟仓配API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/wms(https://pkg.go.dev/github.com/bububa/opentaobao/api/wms)

- [x] 酒店API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/xhotel(https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotel)

- [x] 虾米API(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/xiami(https://pkg.go.dev/github.com/bububa/opentaobao/api/xiami)

- [x] YunOS(https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=0&docType=2&_tb_token_=), github.com/bububa/opentaobao/api/yunos(https://pkg.go.dev/github.com/bububa/opentaobao/api/yunos)


