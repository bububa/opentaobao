# 淘宝开放平台 golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/opentaobao.svg)](https://pkg.go.dev/github.com/bububa/opentaobao)
[![Go](https://github.com/bububa/opentaobao/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/opentaobao/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/opentaobao/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/opentaobao/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/opentaobao.svg)](https://github.com/bububa/opentaobao)


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
| [x] [用户API](https://open.taobao.com/API.htm?docId=24938&docType=2) | [github.com/bububa/opentaobao/api/user](https://pkg.go.dev/github.com/bububa/opentaobao/api/user) |
| [x] [类目API](https://open.taobao.com/API.htm?docId=47659&docType=2) | [github.com/bububa/opentaobao/api/category](https://pkg.go.dev/github.com/bububa/opentaobao/api/category) |
| [x] [商品API](https://open.taobao.com/API.htm?docId=38593&docType=2) | [github.com/bububa/opentaobao/api/product](https://pkg.go.dev/github.com/bububa/opentaobao/api/product) |
| [x] [交易API](https://open.taobao.com/API.htm?docId=48314&docType=2) | [github.com/bububa/opentaobao/api/trade](https://pkg.go.dev/github.com/bububa/opentaobao/api/trade) |
| [x] [评价API](https://open.taobao.com/API.htm?docId=48881&docType=2) | [github.com/bububa/opentaobao/api/traderate](https://pkg.go.dev/github.com/bububa/opentaobao/api/traderate) |
| [x] [物流API](https://open.taobao.com/API.htm?docId=47316&docType=2) | [github.com/bububa/opentaobao/api/logistic](https://pkg.go.dev/github.com/bububa/opentaobao/api/logistic) |
| [x] [店铺API](https://open.taobao.com/API.htm?docId=30898&docType=2) | [github.com/bububa/opentaobao/api/shop](https://pkg.go.dev/github.com/bububa/opentaobao/api/shop) |
| [x] [分销API](https://open.taobao.com/API.htm?docId=55750&docType=2) | [github.com/bububa/opentaobao/api/fenxiao](https://pkg.go.dev/github.com/bububa/opentaobao/api/fenxiao) |
| [x] [旺旺API](https://open.taobao.com/API.htm?docId=27589&docType=2) | [github.com/bububa/opentaobao/api/wangwang](https://pkg.go.dev/github.com/bububa/opentaobao/api/wangwang) |
| [x] [淘宝客API](https://open.taobao.com/API.htm?docId=48340&docType=2) | [github.com/bububa/opentaobao/api/tbk](https://pkg.go.dev/github.com/bububa/opentaobao/api/tbk) |
| [x] [工具API](https://open.taobao.com/API.htm?docId=47640&docType=2) | [github.com/bububa/opentaobao/api/util](https://pkg.go.dev/github.com/bububa/opentaobao/api/util) |
| [x] [物流宝API](https://open.taobao.com/API.htm?docId=28394&docType=2) | [github.com/bububa/opentaobao/api/wlb](https://pkg.go.dev/github.com/bububa/opentaobao/api/wlb) |
| [x] [直通车API](https://open.taobao.com/API.htm?docId=10551&docType=2) | [github.com/bububa/opentaobao/api/simba](https://pkg.go.dev/github.com/bububa/opentaobao/api/simba) |
| [x] [机票API](https://open.taobao.com/API.htm?docId=54268&docType=2) | [github.com/bububa/opentaobao/api/flight](https://pkg.go.dev/github.com/bububa/opentaobao/api/flight) |
| [x] [ONS消息服务](https://open.taobao.com/API.htm?docId=25635&docType=2) | [github.com/bububa/opentaobao/api/jms](https://pkg.go.dev/github.com/bububa/opentaobao/api/jms) |
| [x] [营销API](https://open.taobao.com/API.htm?docId=50923&docType=2) | [github.com/bububa/opentaobao/api/promotion](https://pkg.go.dev/github.com/bububa/opentaobao/api/promotion) |
| [x] [数据API](https://open.taobao.com/API.htm?docId=38780&docType=2) | [github.com/bububa/opentaobao/api/dt](https://pkg.go.dev/github.com/bububa/opentaobao/api/dt) |
| [x] [酒店API](https://open.taobao.com/API.htm?docId=49115&docType=2) | [github.com/bububa/opentaobao/api/xhotel](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotel) |
| [x] [聚划算API](https://open.taobao.com/API.htm?docId=28762&docType=2) | [github.com/bububa/opentaobao/api/ju](https://pkg.go.dev/github.com/bububa/opentaobao/api/ju) |
| [x] [店铺会员管理API](https://open.taobao.com/API.htm?docId=25584&docType=2) | [github.com/bububa/opentaobao/api/crm](https://pkg.go.dev/github.com/bububa/opentaobao/api/crm) |
| [x] [多媒体平台API](https://open.taobao.com/API.htm?docId=27943&docType=2) | [github.com/bububa/opentaobao/api/media](https://pkg.go.dev/github.com/bububa/opentaobao/api/media) |
| [x] [子账号管理API](https://open.taobao.com/API.htm?docId=10820&docType=2) | [github.com/bububa/opentaobao/api/subuser](https://pkg.go.dev/github.com/bububa/opentaobao/api/subuser) |
| [x] [服务平台API](https://open.taobao.com/API.htm?docId=25687&docType=2) | [github.com/bububa/opentaobao/api/servicecenter](https://pkg.go.dev/github.com/bububa/opentaobao/api/servicecenter) |
| [x] [退款API](https://open.taobao.com/API.htm?docId=46784&docType=2) | [github.com/bububa/opentaobao/api/refund](https://pkg.go.dev/github.com/bububa/opentaobao/api/refund) |
| [x] [质检品控API](https://open.taobao.com/API.htm?docId=10902&docType=2) | [github.com/bububa/opentaobao/api/qt](https://pkg.go.dev/github.com/bububa/opentaobao/api/qt) |
| [x] [天猫服务商品API](https://open.taobao.com/API.htm?docId=49044&docType=2) | [github.com/bububa/opentaobao/api/tmallservice](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallservice) |
| [x] [天猫精品库API](https://open.taobao.com/API.htm?docId=21643&docType=2) | [github.com/bububa/opentaobao/api/tmallitem](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallitem) |
| [x] [聚石塔API](https://open.taobao.com/API.htm?docId=55254&docType=2) | [github.com/bububa/opentaobao/api/jst](https://pkg.go.dev/github.com/bububa/opentaobao/api/jst) |
| [x] [电子物流API](https://open.taobao.com/API.htm?docId=27156&docType=2) | [github.com/bububa/opentaobao/api/eticket](https://pkg.go.dev/github.com/bububa/opentaobao/api/eticket) |
| [x] [彩票API](https://open.taobao.com/API.htm?docId=21828&docType=2) | [github.com/bububa/opentaobao/api/caipiao](https://pkg.go.dev/github.com/bububa/opentaobao/api/caipiao) |
| [x] [账务API](https://open.taobao.com/API.htm?docId=21709&docType=2) | [github.com/bububa/opentaobao/api/bill](https://pkg.go.dev/github.com/bububa/opentaobao/api/bill) |
| [x] [拍卖API](https://open.taobao.com/API.htm?docId=38618&docType=2) | [github.com/bububa/opentaobao/api/paimai](https://pkg.go.dev/github.com/bububa/opentaobao/api/paimai) |
| [x] [千牛接口](https://open.taobao.com/API.htm?docId=27244&docType=2) | [github.com/bububa/opentaobao/api/qianniu](https://pkg.go.dev/github.com/bububa/opentaobao/api/qianniu) |
| [x] [消息服务API](https://open.taobao.com/API.htm?docId=53697&docType=2) | [github.com/bububa/opentaobao/api/tmc](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmc) |
| [x] [本地生活API](https://open.taobao.com/API.htm?docId=57060&docType=2) | [github.com/bububa/opentaobao/api/alsc](https://pkg.go.dev/github.com/bububa/opentaobao/api/alsc) |
| [x] [阿里云ocsAPI](https://open.taobao.com/API.htm?docId=24698&docType=2) | [github.com/bububa/opentaobao/api/aliyunocs](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliyunocs) |
| [x] [淘宝同城API](https://open.taobao.com/API.htm?docId=51645&docType=2) | [github.com/bububa/opentaobao/api/cityretail](https://pkg.go.dev/github.com/bububa/opentaobao/api/cityretail) |
| [x] [YunOS](https://open.taobao.com/API.htm?docId=33101&docType=2) | [github.com/bububa/opentaobao/api/yunos](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunos) |
| [x] [阿里云API](https://open.taobao.com/API.htm?docId=24304&docType=2) | [github.com/bububa/opentaobao/api/aliyun](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliyun) |
| [x] [火车票API](https://open.taobao.com/API.htm?docId=24270&docType=2) | [github.com/bububa/opentaobao/api/train](https://pkg.go.dev/github.com/bububa/opentaobao/api/train) |
| [x] [Tanx API](https://open.taobao.com/API.htm?docId=25062&docType=2) | [github.com/bububa/opentaobao/api/tanx](https://pkg.go.dev/github.com/bububa/opentaobao/api/tanx) |
| [x] [手淘开放API](https://open.taobao.com/API.htm?docId=26114&docType=2) | [github.com/bububa/opentaobao/api/mtopopen](https://pkg.go.dev/github.com/bububa/opentaobao/api/mtopopen) |
| [x] [JAE](https://open.taobao.com/API.htm?docId=30901&docType=2) | [github.com/bububa/opentaobao/api/jae](https://pkg.go.dev/github.com/bububa/opentaobao/api/jae) |
| [x] [宝点API](https://open.taobao.com/API.htm?docId=22800&docType=2) | [github.com/bububa/opentaobao/api/baodian](https://pkg.go.dev/github.com/bububa/opentaobao/api/baodian) |
| [x] [天猫会员积分](https://open.taobao.com/API.htm?docId=25758&docType=2) | [github.com/bububa/opentaobao/api/tmallcms](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallcms) |
| [x] [汽车票API](https://open.taobao.com/API.htm?docId=52971&docType=2) | [github.com/bububa/opentaobao/api/bus](https://pkg.go.dev/github.com/bububa/opentaobao/api/bus) |
| [x] [码上淘API](https://open.taobao.com/API.htm?docId=23660&docType=2) | [github.com/bububa/opentaobao/api/ma](https://pkg.go.dev/github.com/bububa/opentaobao/api/ma) |
| [x] [游戏激励平台API](https://open.taobao.com/API.htm?docId=25396&docType=2) | [github.com/bububa/opentaobao/api/gameact](https://pkg.go.dev/github.com/bububa/opentaobao/api/gameact) |
| [x] [淘宝抽奖平台API](https://open.taobao.com/API.htm?docId=23213&docType=2) | [github.com/bububa/opentaobao/api/choujiang](https://pkg.go.dev/github.com/bububa/opentaobao/api/choujiang) |
| [x] [天猫国际API](https://open.taobao.com/API.htm?docId=51012&docType=2) | [github.com/bububa/opentaobao/api/tmallhk](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallhk) |
| [x] [司法拍卖](https://open.taobao.com/API.htm?docId=40551&docType=2) | [github.com/bububa/opentaobao/api/auction](https://pkg.go.dev/github.com/bububa/opentaobao/api/auction) |
| [x] [虾米API](https://open.taobao.com/API.htm?docId=23279&docType=2) | [github.com/bububa/opentaobao/api/xiami](https://pkg.go.dev/github.com/bububa/opentaobao/api/xiami) |
| [x] [天猫互动接口](https://open.taobao.com/API.htm?docId=25768&docType=2) | [github.com/bububa/opentaobao/api/interact](https://pkg.go.dev/github.com/bububa/opentaobao/api/interact) |
| [x] [DMP API](https://open.taobao.com/API.htm?docId=23625&docType=2) | [github.com/bububa/opentaobao/api/dmp](https://pkg.go.dev/github.com/bububa/opentaobao/api/dmp) |
| [x] [生活服务API](https://open.taobao.com/API.htm?docId=50111&docType=2) | [github.com/bububa/opentaobao/api/lifeservice](https://pkg.go.dev/github.com/bububa/opentaobao/api/lifeservice) |
| [x] [手机淘宝API](https://open.taobao.com/API.htm?docId=23431&docType=2) | [github.com/bububa/opentaobao/api/mtop](https://pkg.go.dev/github.com/bububa/opentaobao/api/mtop) |
| [x] [物联API](https://open.taobao.com/API.htm?docId=28313&docType=2) | [github.com/bububa/opentaobao/api/alink](https://pkg.go.dev/github.com/bububa/opentaobao/api/alink) |
| [x] [酒店导购API](https://open.taobao.com/API.htm?docId=35264&docType=2) | [github.com/bububa/opentaobao/api/hotel](https://pkg.go.dev/github.com/bububa/opentaobao/api/hotel) |
| [x] [保险API](https://open.taobao.com/API.htm?docId=37741&docType=2) | [github.com/bububa/opentaobao/api/baoxian](https://pkg.go.dev/github.com/bububa/opentaobao/api/baoxian) |
| [x] [天猫美妆API](https://open.taobao.com/API.htm?docId=25013&docType=2) | [github.com/bububa/opentaobao/api/mei](https://pkg.go.dev/github.com/bububa/opentaobao/api/mei) |
| [x] [会员卡](https://open.taobao.com/API.htm?docId=46278&docType=2) | [github.com/bububa/opentaobao/api/blackvip](https://pkg.go.dev/github.com/bububa/opentaobao/api/blackvip) |
| [x] [电子面单API](https://open.taobao.com/API.htm?docId=27111&docType=2) | [github.com/bububa/opentaobao/api/waybill](https://pkg.go.dev/github.com/bububa/opentaobao/api/waybill) |
| [x] [电影票API](https://open.taobao.com/API.htm?docId=45077&docType=2) | [github.com/bububa/opentaobao/api/film](https://pkg.go.dev/github.com/bububa/opentaobao/api/film) |
| [x] [阿里通信API](https://open.taobao.com/API.htm?docId=29701&docType=2) | [github.com/bububa/opentaobao/api/alicom](https://pkg.go.dev/github.com/bububa/opentaobao/api/alicom) |
| [x] [openimAPI](https://open.taobao.com/API.htm?docId=25766&docType=2) | [github.com/bububa/opentaobao/api/openim](https://pkg.go.dev/github.com/bububa/opentaobao/api/openim) |
| [x] [阿里车联网API](https://open.taobao.com/API.htm?docId=24740&docType=2) | [github.com/bububa/opentaobao/api/autonavi](https://pkg.go.dev/github.com/bububa/opentaobao/api/autonavi) |
| [x] [虚拟院线API](https://open.taobao.com/API.htm?docId=35818&docType=2) | [github.com/bububa/opentaobao/api/taotv](https://pkg.go.dev/github.com/bububa/opentaobao/api/taotv) |
| [x] [知识库API](https://open.taobao.com/API.htm?docId=38771&docType=2) | [github.com/bububa/opentaobao/api/kclub](https://pkg.go.dev/github.com/bububa/opentaobao/api/kclub) |
| [x] [反欺诈风控API](https://open.taobao.com/API.htm?docId=25505&docType=2) | [github.com/bububa/opentaobao/api/antifraud](https://pkg.go.dev/github.com/bububa/opentaobao/api/antifraud) |
| [x] [国际站外贸直通车API](https://open.taobao.com/API.htm?docId=25200&docType=2) | [github.com/bububa/opentaobao/api/scbp](https://pkg.go.dev/github.com/bububa/opentaobao/api/scbp) |
| [x] [天猫服务数据API](https://open.taobao.com/API.htm?docId=53341&docType=2) | [github.com/bububa/opentaobao/api/tmallsc](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallsc) |
| [x] [智能设备](https://open.taobao.com/API.htm?docId=40674&docType=2) | [github.com/bububa/opentaobao/api/iot](https://pkg.go.dev/github.com/bububa/opentaobao/api/iot) |
| [x] [百川](https://open.taobao.com/API.htm?docId=31057&docType=2) | [github.com/bububa/opentaobao/api/baichuan](https://pkg.go.dev/github.com/bububa/opentaobao/api/baichuan) |
| [x] [文本算法API](https://open.taobao.com/API.htm?docId=26130&docType=2) | [github.com/bububa/opentaobao/api/nlp](https://pkg.go.dev/github.com/bububa/opentaobao/api/nlp) |
| [x] [百川推送](https://open.taobao.com/API.htm?docId=25038&docType=2) | [github.com/bububa/opentaobao/api/cloudpush](https://pkg.go.dev/github.com/bububa/opentaobao/api/cloudpush) |
| [x] [国际站商品API](https://open.taobao.com/API.htm?docId=25348&docType=2) | [github.com/bububa/opentaobao/api/icbu](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbu) |
| [x] [淘宝游戏API](https://open.taobao.com/API.htm?docId=45035&docType=2) | [github.com/bububa/opentaobao/api/game](https://pkg.go.dev/github.com/bububa/opentaobao/api/game) |
| [x] [聚安全API](https://open.taobao.com/API.htm?docId=27312&docType=2) | [github.com/bububa/opentaobao/api/security](https://pkg.go.dev/github.com/bububa/opentaobao/api/security) |
| [x] [喵街API](https://open.taobao.com/API.htm?docId=55338&docType=2) | [github.com/bububa/opentaobao/api/mos](https://pkg.go.dev/github.com/bububa/opentaobao/api/mos) |
| [x] [菜鸟配送API](https://open.taobao.com/API.htm?docId=25659&docType=2) | [github.com/bububa/opentaobao/api/cntms](https://pkg.go.dev/github.com/bububa/opentaobao/api/cntms) |
| [x] [菜鸟仓配API](https://open.taobao.com/API.htm?docId=26245&docType=2) | [github.com/bububa/opentaobao/api/wms](https://pkg.go.dev/github.com/bububa/opentaobao/api/wms) |
| [x] [网上法庭对外API](https://open.taobao.com/API.htm?docId=27772&docType=2) | [github.com/bububa/opentaobao/api/nazca](https://pkg.go.dev/github.com/bububa/opentaobao/api/nazca) |
| [x] [五道口API](https://open.taobao.com/API.htm?docId=53548&docType=2) | [github.com/bububa/opentaobao/api/wdk](https://pkg.go.dev/github.com/bububa/opentaobao/api/wdk) |
| [x] [阿里大于API](https://open.taobao.com/API.htm?docId=28647&docType=2) | [github.com/bububa/opentaobao/api/aliqin](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliqin) |
| [x] [淘宝内容API](https://open.taobao.com/API.htm?docId=27196&docType=2) | [github.com/bububa/opentaobao/api/beehive](https://pkg.go.dev/github.com/bububa/opentaobao/api/beehive) |
| [x] [旅行用车API](https://open.taobao.com/API.htm?docId=44234&docType=2) | [github.com/bububa/opentaobao/api/car](https://pkg.go.dev/github.com/bububa/opentaobao/api/car) |
| [x] [门票-商品管理API](https://open.taobao.com/API.htm?docId=27945&docType=2) | [github.com/bububa/opentaobao/api/ticket](https://pkg.go.dev/github.com/bububa/opentaobao/api/ticket) |
| [x] [菜鸟无线API](https://open.taobao.com/API.htm?docId=26056&docType=2) | [github.com/bububa/opentaobao/api/guoguo](https://pkg.go.dev/github.com/bububa/opentaobao/api/guoguo) |
| [x] [奇门仓储API](https://open.taobao.com/API.htm?docId=29364&docType=2) | [github.com/bububa/opentaobao/api/qimen](https://pkg.go.dev/github.com/bububa/opentaobao/api/qimen) |
| [x] [yunos推送服务api](https://open.taobao.com/API.htm?docId=49740&docType=2) | [github.com/bububa/opentaobao/api/cmns](https://pkg.go.dev/github.com/bububa/opentaobao/api/cmns) |
| [x] [生活汇API](https://open.taobao.com/API.htm?docId=26189&docType=2) | [github.com/bububa/opentaobao/api/elife](https://pkg.go.dev/github.com/bububa/opentaobao/api/elife) |
| [x] [tv支付](https://open.taobao.com/API.htm?docId=26205&docType=2) | [github.com/bububa/opentaobao/api/tvpay](https://pkg.go.dev/github.com/bububa/opentaobao/api/tvpay) |
| [x] [菜鸟集货API](https://open.taobao.com/API.htm?docId=45980&docType=2) | [github.com/bububa/opentaobao/api/wlbimports](https://pkg.go.dev/github.com/bububa/opentaobao/api/wlbimports) |
| [x] [阿里健康医](https://open.taobao.com/API.htm?docId=42645&docType=2) | [github.com/bububa/opentaobao/api/alihealth](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealth) |
| [x] [地动仪](https://open.taobao.com/API.htm?docId=37287&docType=2) | [github.com/bububa/opentaobao/api/lbs](https://pkg.go.dev/github.com/bububa/opentaobao/api/lbs) |
| [x] [阿里健康药API](https://open.taobao.com/API.htm?docId=50465&docType=2) | [github.com/bububa/opentaobao/api/drug](https://pkg.go.dev/github.com/bububa/opentaobao/api/drug) |
| [x] [手淘分享](https://open.taobao.com/API.htm?docId=32461&docType=2) | [github.com/bububa/opentaobao/api/wirelessshare](https://pkg.go.dev/github.com/bububa/opentaobao/api/wirelessshare) |
| [x] [法务诉讼对外API](https://open.taobao.com/API.htm?docId=53340&docType=2) | [github.com/bububa/opentaobao/api/legalcase](https://pkg.go.dev/github.com/bububa/opentaobao/api/legalcase) |
| [x] [度假-商品管理API](https://open.taobao.com/API.htm?docId=28132&docType=2) | [github.com/bububa/opentaobao/api/travel](https://pkg.go.dev/github.com/bububa/opentaobao/api/travel) |
| [x] [酒店商品API](https://open.taobao.com/API.htm?docId=22665&docType=2) | [github.com/bububa/opentaobao/api/xhotelitem](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotelitem) |
| [x] [酒店在线预订API](https://open.taobao.com/API.htm?docId=48716&docType=2) | [github.com/bububa/opentaobao/api/xhotelonlineorder](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotelonlineorder) |
| [x] [酒店官网信用住API](https://open.taobao.com/API.htm?docId=26073&docType=2) | [github.com/bububa/opentaobao/api/xhotelofficial](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotelofficial) |
| [x] [酒店线下信用住API](https://open.taobao.com/API.htm?docId=26074&docType=2) | [github.com/bububa/opentaobao/api/xhoteloffline](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhoteloffline) |
| [x] [全渠道API](https://open.taobao.com/API.htm?docId=42927&docType=2) | [github.com/bububa/opentaobao/api/omniorder](https://pkg.go.dev/github.com/bububa/opentaobao/api/omniorder) |
| [x] [国际机票政策API](https://open.taobao.com/API.htm?docId=25494&docType=2) | [github.com/bububa/opentaobao/api/itpolicy](https://pkg.go.dev/github.com/bububa/opentaobao/api/itpolicy) |
| [x] [国际机票订单API](https://open.taobao.com/API.htm?docId=27905&docType=2) | [github.com/bububa/opentaobao/api/ieagency](https://pkg.go.dev/github.com/bububa/opentaobao/api/ieagency) |
| [x] [国内机票订单API](https://open.taobao.com/API.htm?docId=26565&docType=2) | [github.com/bububa/opentaobao/api/jipiao](https://pkg.go.dev/github.com/bububa/opentaobao/api/jipiao) |
| [x] [度假&门票-交易管理API](https://open.taobao.com/API.htm?docId=52211&docType=2) | [github.com/bububa/opentaobao/api/traveltrade](https://pkg.go.dev/github.com/bububa/opentaobao/api/traveltrade) |
| [x] [阿里体育API](https://open.taobao.com/API.htm?docId=41267&docType=2) | [github.com/bububa/opentaobao/api/alisports](https://pkg.go.dev/github.com/bububa/opentaobao/api/alisports) |
| [x] [电子发票](https://open.taobao.com/API.htm?docId=27764&docType=2) | [github.com/bububa/opentaobao/api/einvoice](https://pkg.go.dev/github.com/bububa/opentaobao/api/einvoice) |
| [x] [阿里翻译API](https://open.taobao.com/API.htm?docId=51871&docType=2) | [github.com/bububa/opentaobao/api/seaking](https://pkg.go.dev/github.com/bububa/opentaobao/api/seaking) |
| [x] [国际站数据管](https://open.taobao.com/API.htm?docId=26908&docType=2) | [github.com/bububa/opentaobao/api/mydata](https://pkg.go.dev/github.com/bububa/opentaobao/api/mydata) |
| [x] [tmall-carcenter](https://open.taobao.com/API.htm?docId=31209&docType=2) | [github.com/bububa/opentaobao/api/tmallcarenter](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallcarenter) |
| [x] [阿里妈妈广告中心API](https://open.taobao.com/API.htm?docId=47165&docType=2) | [github.com/bububa/opentaobao/api/scs](https://pkg.go.dev/github.com/bububa/opentaobao/api/scs) |
| [x] [Yunos Account API](https://open.taobao.com/API.htm?docId=27319&docType=2) | [github.com/bububa/opentaobao/api/yunosaccount](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunosaccount) |
| [x] [菜鸟裹裹API](https://open.taobao.com/API.htm?docId=31197&docType=2) | [github.com/bububa/opentaobao/api/cainiaolocker](https://pkg.go.dev/github.com/bububa/opentaobao/api/cainiaolocker) |
| [x] [度假-签证管理API](https://open.taobao.com/API.htm?docId=44453&docType=2) | [github.com/bububa/opentaobao/api/normalvisa](https://pkg.go.dev/github.com/bububa/opentaobao/api/normalvisa) |
| [x] [菜鸟农村物流](https://open.taobao.com/API.htm?docId=27588&docType=2) | [github.com/bububa/opentaobao/api/cainiaoncwl](https://pkg.go.dev/github.com/bububa/opentaobao/api/cainiaoncwl) |
| [x] [体检机构API](https://open.taobao.com/API.htm?docId=41650&docType=2) | [github.com/bububa/opentaobao/api/examination](https://pkg.go.dev/github.com/bububa/opentaobao/api/examination) |
| [x] [1688推客API](https://open.taobao.com/API.htm?docId=27457&docType=2) | [github.com/bububa/opentaobao/api/tuike](https://pkg.go.dev/github.com/bububa/opentaobao/api/tuike) |
| [x] [商户API](https://open.taobao.com/API.htm?docId=50163&docType=2) | [github.com/bububa/opentaobao/api/store](https://pkg.go.dev/github.com/bububa/opentaobao/api/store) |
| [x] [桌面API](https://open.taobao.com/API.htm?docId=29910&docType=2) | [github.com/bububa/opentaobao/api/ott](https://pkg.go.dev/github.com/bububa/opentaobao/api/ott) |
| [x] [电动车API](https://open.taobao.com/API.htm?docId=30434&docType=2) | [github.com/bububa/opentaobao/api/vms](https://pkg.go.dev/github.com/bububa/opentaobao/api/vms) |
| [x] [新零售API](https://open.taobao.com/API.htm?docId=44781&docType=2) | [github.com/bububa/opentaobao/api/newretail](https://pkg.go.dev/github.com/bububa/opentaobao/api/newretail) |
| [x] [SCM API](https://open.taobao.com/API.htm?docId=36216&docType=2) | [github.com/bububa/opentaobao/api/ascm](https://pkg.go.dev/github.com/bububa/opentaobao/api/ascm) |
| [x] [百川-ctg](https://open.taobao.com/API.htm?docId=28682&docType=2) | [github.com/bububa/opentaobao/api/baichuanctg](https://pkg.go.dev/github.com/bububa/opentaobao/api/baichuanctg) |
| [x] [汇金API](https://open.taobao.com/API.htm?docId=30278&docType=2) | [github.com/bububa/opentaobao/api/fundplatform](https://pkg.go.dev/github.com/bububa/opentaobao/api/fundplatform) |
| [x] [数娱媒资输出](https://open.taobao.com/API.htm?docId=41681&docType=2) | [github.com/bububa/opentaobao/api/wenyuvideo](https://pkg.go.dev/github.com/bububa/opentaobao/api/wenyuvideo) |
| [x] [商旅API](https://open.taobao.com/API.htm?docId=33169&docType=2) | [github.com/bububa/opentaobao/api/btrip](https://pkg.go.dev/github.com/bububa/opentaobao/api/btrip) |
| [x] [零售plus](https://open.taobao.com/API.htm?docId=36762&docType=2) | [github.com/bububa/opentaobao/api/nlife](https://pkg.go.dev/github.com/bububa/opentaobao/api/nlife) |
| [x] [天猫门店API](https://open.taobao.com/API.htm?docId=31454&docType=2) | [github.com/bububa/opentaobao/api/retail](https://pkg.go.dev/github.com/bububa/opentaobao/api/retail) |
| [x] [AILAB图像算法API](https://open.taobao.com/API.htm?docId=37057&docType=2) | [github.com/bububa/opentaobao/api/aiar](https://pkg.go.dev/github.com/bububa/opentaobao/api/aiar) |
| [x] [天猫供应](https://open.taobao.com/API.htm?docId=40826&docType=2) | [github.com/bububa/opentaobao/api/ascp](https://pkg.go.dev/github.com/bububa/opentaobao/api/ascp) |
| [x] [欢行开发平台API](https://open.taobao.com/API.htm?docId=47501&docType=2) | [github.com/bububa/opentaobao/api/happytrip](https://pkg.go.dev/github.com/bububa/opentaobao/api/happytrip) |
| [x] [跨境API](https://open.taobao.com/API.htm?docId=28703&docType=2) | [github.com/bububa/opentaobao/api/oversea](https://pkg.go.dev/github.com/bububa/opentaobao/api/oversea) |
| [x] [迎客松牌照审核接口](https://open.taobao.com/API.htm?docId=30410&docType=2) | [github.com/bububa/opentaobao/api/tvupadmin](https://pkg.go.dev/github.com/bububa/opentaobao/api/tvupadmin) |
| [x] [IoTI API](https://open.taobao.com/API.htm?docId=46348&docType=2) | [github.com/bububa/opentaobao/api/ioti](https://pkg.go.dev/github.com/bububa/opentaobao/api/ioti) |
| [x] [淘宝卡券平台](https://open.taobao.com/API.htm?docId=29874&docType=2) | [github.com/bububa/opentaobao/api/deliveryvoucher](https://pkg.go.dev/github.com/bububa/opentaobao/api/deliveryvoucher) |
| [x] [智慧门店](https://open.taobao.com/API.htm?docId=39618&docType=2) | [github.com/bububa/opentaobao/api/smartstore](https://pkg.go.dev/github.com/bububa/opentaobao/api/smartstore) |
| [x] [淘宝定制行业API](https://open.taobao.com/API.htm?docId=28671&docType=2) | [github.com/bububa/opentaobao/api/customizemarket](https://pkg.go.dev/github.com/bububa/opentaobao/api/customizemarket) |
| [x] [闲鱼](https://open.taobao.com/API.htm?docId=48312&docType=2) | [github.com/bububa/opentaobao/api/idle](https://pkg.go.dev/github.com/bububa/opentaobao/api/idle) |
| [x] [库存API](https://open.taobao.com/API.htm?docId=31754&docType=2) | [github.com/bububa/opentaobao/api/inventory](https://pkg.go.dev/github.com/bububa/opentaobao/api/inventory) |
| [x] [aliExpress](https://open.taobao.com/API.htm?docId=54728&docType=2) | [github.com/bububa/opentaobao/api/aliexpress](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliexpress) |
| [x] [YunOS-广告](https://open.taobao.com/API.htm?docId=29086&docType=2) | [github.com/bububa/opentaobao/api/yunosad](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunosad) |
| [x] [司法开放平台](https://open.taobao.com/API.htm?docId=43036&docType=2) | [github.com/bububa/opentaobao/api/legalsuit](https://pkg.go.dev/github.com/bububa/opentaobao/api/legalsuit) |
| [x] [全球速卖通](https://open.taobao.com/API.htm?docId=56229&docType=2) | [github.com/bububa/opentaobao/api/aliexpresssumaitong](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliexpresssumaitong) |
| [x] [阿里健康追溯码](https://open.taobao.com/API.htm?docId=40971&docType=2) | [github.com/bububa/opentaobao/api/drugtrace](https://pkg.go.dev/github.com/bububa/opentaobao/api/drugtrace) |
| [x] [会员中心API](https://open.taobao.com/API.htm?docId=29108&docType=2) | [github.com/bububa/opentaobao/api/interactvip](https://pkg.go.dev/github.com/bububa/opentaobao/api/interactvip) |
| [x] [阿里健康-会员管理](https://open.taobao.com/API.htm?docId=46804&docType=2) | [github.com/bububa/opentaobao/api/alihealthcrm](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthcrm) |
| [x] [品效API](https://open.taobao.com/API.htm?docId=36460&docType=2) | [github.com/bububa/opentaobao/api/brandhub](https://pkg.go.dev/github.com/bububa/opentaobao/api/brandhub) |
| [x] [小蜜API](https://open.taobao.com/API.htm?docId=44065&docType=2) | [github.com/bububa/opentaobao/api/alime](https://pkg.go.dev/github.com/bububa/opentaobao/api/alime) |
| [x] [大麦票务云分销API](https://open.taobao.com/API.htm?docId=49119&docType=2) | [github.com/bububa/opentaobao/api/maitix](https://pkg.go.dev/github.com/bububa/opentaobao/api/maitix) |
| [x] [天猫精灵开放API](https://open.taobao.com/API.htm?docId=36465&docType=2) | [github.com/bububa/opentaobao/api/tmallgenie](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallgenie) |
| [x] [ICBU卖家API](https://open.taobao.com/API.htm?docId=51234&docType=2) | [github.com/bububa/opentaobao/api/icbuseller](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbuseller) |
| [x] [智慧园区API](https://open.taobao.com/API.htm?docId=32226&docType=2) | [github.com/bububa/opentaobao/api/campus](https://pkg.go.dev/github.com/bububa/opentaobao/api/campus) |
| [x] [酒店会员API](https://open.taobao.com/API.htm?docId=34020&docType=2) | [github.com/bububa/opentaobao/api/xhotelcrm](https://pkg.go.dev/github.com/bububa/opentaobao/api/xhotelcrm) |
| [x] [大麦API](https://open.taobao.com/API.htm?docId=39475&docType=2) | [github.com/bububa/opentaobao/api/damai](https://pkg.go.dev/github.com/bububa/opentaobao/api/damai) |
| [x] [换货API](https://open.taobao.com/API.htm?docId=31202&docType=2) | [github.com/bububa/opentaobao/api/exchange](https://pkg.go.dev/github.com/bububa/opentaobao/api/exchange) |
| [x] [商家营销中心API](https://open.taobao.com/API.htm?docId=31352&docType=2) | [github.com/bububa/opentaobao/api/singletreasure](https://pkg.go.dev/github.com/bububa/opentaobao/api/singletreasure) |
| [x] [ICBU－物流](https://open.taobao.com/API.htm?docId=47959&docType=2) | [github.com/bububa/opentaobao/api/icbulogistics](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbulogistics) |
| [x] [ICBU－RFQ](https://open.taobao.com/API.htm?docId=40805&docType=2) | [github.com/bububa/opentaobao/api/icburfq](https://pkg.go.dev/github.com/bububa/opentaobao/api/icburfq) |
| [x] [ICBU－信保](https://open.taobao.com/API.htm?docId=31686&docType=2) | [github.com/bububa/opentaobao/api/icbuassurance](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbuassurance) |
| [x] [飞猪POI数据API](https://open.taobao.com/API.htm?docId=45315&docType=2) | [github.com/bububa/opentaobao/api/alitrippoi](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitrippoi) |
| [x] [飞猪行政区划API](https://open.taobao.com/API.htm?docId=31375&docType=2) | [github.com/bububa/opentaobao/api/alitripdivisions](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitripdivisions) |
| [x] [零售终端API](https://open.taobao.com/API.htm?docId=35424&docType=2) | [github.com/bububa/opentaobao/api/nrt](https://pkg.go.dev/github.com/bububa/opentaobao/api/nrt) |
| [x] [手淘用户增](https://open.taobao.com/API.htm?docId=43959&docType=2) | [github.com/bububa/opentaobao/api/usergrowth](https://pkg.go.dev/github.com/bububa/opentaobao/api/usergrowth) |
| [x] [ALiOS应用中心](https://open.taobao.com/API.htm?docId=35428&docType=2) | [github.com/bububa/opentaobao/api/yunosappstore](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunosappstore) |
| [x] [b2b认证平台api](https://open.taobao.com/API.htm?docId=32288&docType=2) | [github.com/bububa/opentaobao/api/b2bcert](https://pkg.go.dev/github.com/bububa/opentaobao/api/b2bcert) |
| [x] [企业运营平台-集团财](https://open.taobao.com/API.htm?docId=33937&docType=2) | [github.com/bububa/opentaobao/api/fpm](https://pkg.go.dev/github.com/bububa/opentaobao/api/fpm) |
| [x] [零售通智能POS开放](https://open.taobao.com/API.htm?docId=32945&docType=2) | [github.com/bububa/opentaobao/api/lstpos](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstpos) |
| [x] [阿里巴巴供应链平台API](https://open.taobao.com/API.htm?docId=32779&docType=2) | [github.com/bububa/opentaobao/api/ascpqcc](https://pkg.go.dev/github.com/bububa/opentaobao/api/ascpqcc) |
| [x] [天猫新零售](https://open.taobao.com/API.htm?docId=55242&docType=2) | [github.com/bububa/opentaobao/api/tmallnr](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallnr) |
| [x] [酒店团购套餐API](https://open.taobao.com/API.htm?docId=45432&docType=2) | [github.com/bububa/opentaobao/api/tuanhotel](https://pkg.go.dev/github.com/bububa/opentaobao/api/tuanhotel) |
| [x] [阿里影业云智API](https://open.taobao.com/API.htm?docId=42042&docType=2) | [github.com/bububa/opentaobao/api/larkiot](https://pkg.go.dev/github.com/bububa/opentaobao/api/larkiot) |
| [x] [渠道中心API](https://open.taobao.com/API.htm?docId=40706&docType=2) | [github.com/bububa/opentaobao/api/tmallchannel](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallchannel) |
| [x] [阿里健康-疫苗API](https://open.taobao.com/API.htm?docId=37710&docType=2) | [github.com/bububa/opentaobao/api/vaccin](https://pkg.go.dev/github.com/bububa/opentaobao/api/vaccin) |
| [x] [教育账号 API](https://open.taobao.com/API.htm?docId=35502&docType=2) | [github.com/bububa/opentaobao/api/yunosdm](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunosdm) |
| [x] [用户增长](https://open.taobao.com/API.htm?docId=43098&docType=2) | [github.com/bububa/opentaobao/api/usergrowth2](https://pkg.go.dev/github.com/bububa/opentaobao/api/usergrowth2) |
| [x] [AE-供应链](https://open.taobao.com/API.htm?docId=56317&docType=2) | [github.com/bububa/opentaobao/api/ascpffo](https://pkg.go.dev/github.com/bububa/opentaobao/api/ascpffo) |
| [x] [互动吧API](https://open.taobao.com/API.htm?docId=44032&docType=2) | [github.com/bububa/opentaobao/api/fans](https://pkg.go.dev/github.com/bububa/opentaobao/api/fans) |
| [x] [飞猪机票前台类目](https://open.taobao.com/API.htm?docId=42575&docType=2) | [github.com/bububa/opentaobao/api/flightuppc](https://pkg.go.dev/github.com/bububa/opentaobao/api/flightuppc) |
| [x] [平台治理API](https://open.taobao.com/API.htm?docId=32353&docType=2) | [github.com/bububa/opentaobao/api/sungari](https://pkg.go.dev/github.com/bububa/opentaobao/api/sungari) |
| [x] [天猫超市前台API](https://open.taobao.com/API.htm?docId=45441&docType=2) | [github.com/bububa/opentaobao/api/txcs](https://pkg.go.dev/github.com/bububa/opentaobao/api/txcs) |
| [x] [虾米开放平台](https://open.taobao.com/API.htm?docId=36051&docType=2) | [github.com/bububa/opentaobao/api/xiamiopen](https://pkg.go.dev/github.com/bububa/opentaobao/api/xiamiopen) |
| [x] [阿里健康医生](https://open.taobao.com/API.htm?docId=43570&docType=2) | [github.com/bububa/opentaobao/api/alihealthoutflow](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthoutflow) |
| [x] [飞猪酒店标准库](https://open.taobao.com/API.htm?docId=44900&docType=2) | [github.com/bububa/opentaobao/api/hotelhstdf](https://pkg.go.dev/github.com/bububa/opentaobao/api/hotelhstdf) |
| [x] [AE-Dropshipper](https://open.taobao.com/API.htm?docId=39371&docType=2) | [github.com/bububa/opentaobao/api/aedropshiper](https://pkg.go.dev/github.com/bububa/opentaobao/api/aedropshiper) |
| [x] [银泰scm-openapi](https://open.taobao.com/API.htm?docId=41072&docType=2) | [github.com/bububa/opentaobao/api/moscm](https://pkg.go.dev/github.com/bububa/opentaobao/api/moscm) |
| [x] [信息平台-采购](https://open.taobao.com/API.htm?docId=38501&docType=2) | [github.com/bububa/opentaobao/api/pur](https://pkg.go.dev/github.com/bububa/opentaobao/api/pur) |
| [x] [零售通自动售货机](https://open.taobao.com/API.htm?docId=37506&docType=2) | [github.com/bububa/opentaobao/api/lstvending](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstvending) |
| [x] [淘宝小程序API](https://open.taobao.com/API.htm?docId=54266&docType=2) | [github.com/bububa/opentaobao/api/miniapp](https://pkg.go.dev/github.com/bububa/opentaobao/api/miniapp) |
| [x] [ott支付](https://open.taobao.com/API.htm?docId=46647&docType=2) | [github.com/bububa/opentaobao/api/ottpay](https://pkg.go.dev/github.com/bububa/opentaobao/api/ottpay) |
| [x] [天猫汽车](https://open.taobao.com/API.htm?docId=41631&docType=2) | [github.com/bububa/opentaobao/api/tmallcar](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallcar) |
| [x] [ALIOS广告平台](https://open.taobao.com/API.htm?docId=37669&docType=2) | [github.com/bububa/opentaobao/api/admarket](https://pkg.go.dev/github.com/bububa/opentaobao/api/admarket) |
| [x] [业务平台新零售](https://open.taobao.com/API.htm?docId=41646&docType=2) | [github.com/bububa/opentaobao/api/uscesl](https://pkg.go.dev/github.com/bububa/opentaobao/api/uscesl) |
| [x] [大麦第三方商家接入API](https://open.taobao.com/API.htm?docId=55148&docType=2) | [github.com/bububa/opentaobao/api/damaiticklet](https://pkg.go.dev/github.com/bububa/opentaobao/api/damaiticklet) |
| [x] [ICBU商品api](https://open.taobao.com/API.htm?docId=50558&docType=2) | [github.com/bububa/opentaobao/api/icbuproduct](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbuproduct) |
| [x] [神鲸应用API](https://open.taobao.com/API.htm?docId=45850&docType=2) | [github.com/bububa/opentaobao/api/shenjing](https://pkg.go.dev/github.com/bububa/opentaobao/api/shenjing) |
| [x] [阿里影业灯塔](https://open.taobao.com/API.htm?docId=50713&docType=2) | [github.com/bububa/opentaobao/api/dengta](https://pkg.go.dev/github.com/bububa/opentaobao/api/dengta) |
| [x] [激励API](https://open.taobao.com/API.htm?docId=38249&docType=2) | [github.com/bububa/opentaobao/api/degoperation](https://pkg.go.dev/github.com/bububa/opentaobao/api/degoperation) |
| [x] [优酷-媒资](https://open.taobao.com/API.htm?docId=42057&docType=2) | [github.com/bububa/opentaobao/api/youkuott](https://pkg.go.dev/github.com/bububa/opentaobao/api/youkuott) |
| [x] [优酷网盟](https://open.taobao.com/API.htm?docId=39065&docType=2) | [github.com/bububa/opentaobao/api/youkudsp](https://pkg.go.dev/github.com/bububa/opentaobao/api/youkudsp) |
| [x] [阿里健康API](https://open.taobao.com/API.htm?docId=43381&docType=2) | [github.com/bububa/opentaobao/api/alihealth2](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealth2) |
| [x] [新制造API](https://open.taobao.com/API.htm?docId=39534&docType=2) | [github.com/bububa/opentaobao/api/rhino](https://pkg.go.dev/github.com/bububa/opentaobao/api/rhino) |
| [x] [零售通小店智能设备](https://open.taobao.com/API.htm?docId=45361&docType=2) | [github.com/bububa/opentaobao/api/lstspeacker](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstspeacker) |
| [x] [医知鹿-视频](https://open.taobao.com/API.htm?docId=46849&docType=2) | [github.com/bububa/opentaobao/api/alihealthmdeer](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthmdeer) |
| [x] [五道口商品API](https://open.taobao.com/API.htm?docId=30648&docType=2) | [github.com/bububa/opentaobao/api/wdkitem](https://pkg.go.dev/github.com/bububa/opentaobao/api/wdkitem) |
| [x] [零售通订单履行](https://open.taobao.com/API.htm?docId=53484&docType=2) | [github.com/bububa/opentaobao/api/lstlogistics](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstlogistics) |
| [x] [天猫线下大屏](https://open.taobao.com/API.htm?docId=39240&docType=2) | [github.com/bububa/opentaobao/api/tmallfcbox](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallfcbox) |
| [x] [AE-Oversea-Solution](https://open.taobao.com/API.htm?docId=47378&docType=2) | [github.com/bububa/opentaobao/api/aesolution](https://pkg.go.dev/github.com/bububa/opentaobao/api/aesolution) |
| [x] [阿里精灵基础能力接口](https://open.taobao.com/API.htm?docId=39668&docType=2) | [github.com/bububa/opentaobao/api/aligenie](https://pkg.go.dev/github.com/bububa/opentaobao/api/aligenie) |
| [x] [阿里健康新零售](https://open.taobao.com/API.htm?docId=41387&docType=2) | [github.com/bububa/opentaobao/api/healthnr](https://pkg.go.dev/github.com/bububa/opentaobao/api/healthnr) |
| [x] [闲鱼发布](https://open.taobao.com/API.htm?docId=39868&docType=2) | [github.com/bububa/opentaobao/api/idleitem](https://pkg.go.dev/github.com/bububa/opentaobao/api/idleitem) |
| [x] [飞猪酒店签约中心](https://open.taobao.com/API.htm?docId=49754&docType=2) | [github.com/bububa/opentaobao/api/hotelalliance](https://pkg.go.dev/github.com/bububa/opentaobao/api/hotelalliance) |
| [x] [人工智能实验室开放平台API](https://open.taobao.com/API.htm?docId=45531&docType=2) | [github.com/bububa/opentaobao/api/alilabs](https://pkg.go.dev/github.com/bububa/opentaobao/api/alilabs) |
| [x] [天猫新品创新中心API](https://open.taobao.com/API.htm?docId=43329&docType=2) | [github.com/bububa/opentaobao/api/tmic](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmic) |
| [x] [MOZI账号API](https://open.taobao.com/API.htm?docId=45946&docType=2) | [github.com/bububa/opentaobao/api/mozi](https://pkg.go.dev/github.com/bububa/opentaobao/api/mozi) |
| [x] [亲橙里westcrmAPI](https://open.taobao.com/API.htm?docId=41384&docType=2) | [github.com/bububa/opentaobao/api/westcrm](https://pkg.go.dev/github.com/bububa/opentaobao/api/westcrm) |
| [x] [ICBU-橱](https://open.taobao.com/API.htm?docId=40964&docType=2) | [github.com/bububa/opentaobao/api/icbushowcase](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbushowcase) |
| [x] [云码API](https://open.taobao.com/API.htm?docId=47156&docType=2) | [github.com/bububa/opentaobao/api/mc](https://pkg.go.dev/github.com/bububa/opentaobao/api/mc) |
| [x] [HOMEAI](https://open.taobao.com/API.htm?docId=41831&docType=2) | [github.com/bububa/opentaobao/api/homeai](https://pkg.go.dev/github.com/bububa/opentaobao/api/homeai) |
| [x] [MOZI权限API](https://open.taobao.com/API.htm?docId=45484&docType=2) | [github.com/bububa/opentaobao/api/moziacl](https://pkg.go.dev/github.com/bububa/opentaobao/api/moziacl) |
| [x] [飞猪商家平台](https://open.taobao.com/API.htm?docId=54748&docType=2) | [github.com/bububa/opentaobao/api/alitripmerchant](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitripmerchant) |
| [x] [新零售POS](https://open.taobao.com/API.htm?docId=42992&docType=2) | [github.com/bububa/opentaobao/api/nrpos](https://pkg.go.dev/github.com/bububa/opentaobao/api/nrpos) |
| [x] [资质共享API](https://open.taobao.com/API.htm?docId=42262&docType=2) | [github.com/bububa/opentaobao/api/fivee](https://pkg.go.dev/github.com/bububa/opentaobao/api/fivee) |
| [x] [业务平台事业部-税务平台API](https://open.taobao.com/API.htm?docId=41740&docType=2) | [github.com/bububa/opentaobao/api/tax](https://pkg.go.dev/github.com/bububa/opentaobao/api/tax) |
| [x] [MOZI 租户](https://open.taobao.com/API.htm?docId=51055&docType=2) | [github.com/bububa/opentaobao/api/mozivds](https://pkg.go.dev/github.com/bububa/opentaobao/api/mozivds) |
| [x] [公益三小时公共](https://open.taobao.com/API.htm?docId=45562&docType=2) | [github.com/bububa/opentaobao/api/charity](https://pkg.go.dev/github.com/bububa/opentaobao/api/charity) |
| [x] [脚模API](https://open.taobao.com/API.htm?docId=50512&docType=2) | [github.com/bububa/opentaobao/api/foodscan](https://pkg.go.dev/github.com/bububa/opentaobao/api/foodscan) |
| [x] [阿里健康处方药平台](https://open.taobao.com/API.htm?docId=44810&docType=2) | [github.com/bububa/opentaobao/api/alidoc](https://pkg.go.dev/github.com/bububa/opentaobao/api/alidoc) |
| [x] [优酷播控幻影API](https://open.taobao.com/API.htm?docId=43082&docType=2) | [github.com/bububa/opentaobao/api/mirage](https://pkg.go.dev/github.com/bububa/opentaobao/api/mirage) |
| [x] [miniapp开放API](https://open.taobao.com/API.htm?docId=51768&docType=2) | [github.com/bububa/opentaobao/api/miniappopen](https://pkg.go.dev/github.com/bububa/opentaobao/api/miniappopen) |
| [x] [天天特价](https://open.taobao.com/API.htm?docId=48644&docType=2) | [github.com/bububa/opentaobao/api/tttm](https://pkg.go.dev/github.com/bububa/opentaobao/api/tttm) |
| [x] [iHome API](https://open.taobao.com/API.htm?docId=43219&docType=2) | [github.com/bububa/opentaobao/api/ihome](https://pkg.go.dev/github.com/bububa/opentaobao/api/ihome) |
| [x] [信息流API](https://open.taobao.com/API.htm?docId=43248&docType=2) | [github.com/bububa/opentaobao/api/feedflow](https://pkg.go.dev/github.com/bububa/opentaobao/api/feedflow) |
| [x] [淘宝C2M](https://open.taobao.com/API.htm?docId=51314&docType=2) | [github.com/bububa/opentaobao/api/c2m](https://pkg.go.dev/github.com/bububa/opentaobao/api/c2m) |
| [x] [新零售供应链API](https://open.taobao.com/API.htm?docId=54020&docType=2) | [github.com/bububa/opentaobao/api/ascpchannel](https://pkg.go.dev/github.com/bububa/opentaobao/api/ascpchannel) |
| [x] [AliOS支付API](https://open.taobao.com/API.htm?docId=45261&docType=2) | [github.com/bububa/opentaobao/api/aliospay](https://pkg.go.dev/github.com/bububa/opentaobao/api/aliospay) |
| [x] [菜鸟控制塔API](https://open.taobao.com/API.htm?docId=43854&docType=2) | [github.com/bububa/opentaobao/api/cainiaoecc](https://pkg.go.dev/github.com/bububa/opentaobao/api/cainiaoecc) |
| [x] [小程序API](https://open.taobao.com/API.htm?docId=54992&docType=2) | [github.com/bububa/opentaobao/api/yunosminiapp](https://pkg.go.dev/github.com/bububa/opentaobao/api/yunosminiapp) |
| [x] [菜鸟末端商业](https://open.taobao.com/API.htm?docId=46493&docType=2) | [github.com/bububa/opentaobao/api/cainiaocntec](https://pkg.go.dev/github.com/bububa/opentaobao/api/cainiaocntec) |
| [x] [阿里健康挂号号源直连](https://open.taobao.com/API.htm?docId=55216&docType=2) | [github.com/bububa/opentaobao/api/medicalbase](https://pkg.go.dev/github.com/bububa/opentaobao/api/medicalbase) |
| [x] [AE-UserOpen-Recommend](https://open.taobao.com/API.htm?docId=45722&docType=2) | [github.com/bububa/opentaobao/api/aeusergrowth](https://pkg.go.dev/github.com/bububa/opentaobao/api/aeusergrowth) |
| [x] [Efficient Tools](https://open.taobao.com/API.htm?docId=45804&docType=2) | [github.com/bububa/opentaobao/api/aetools](https://pkg.go.dev/github.com/bububa/opentaobao/api/aetools) |
| [x] [Data Reports](https://open.taobao.com/API.htm?docId=45807&docType=2) | [github.com/bububa/opentaobao/api/aedata](https://pkg.go.dev/github.com/bububa/opentaobao/api/aedata) |
| [x] [Promotion Creatives](https://open.taobao.com/API.htm?docId=45801&docType=2) | [github.com/bububa/opentaobao/api/aecreatives](https://pkg.go.dev/github.com/bububa/opentaobao/api/aecreatives) |
| [x] [菜鸟国际出口](https://open.taobao.com/API.htm?docId=46429&docType=2) | [github.com/bububa/opentaobao/api/cainiaohandover](https://pkg.go.dev/github.com/bububa/opentaobao/api/cainiaohandover) |
| [x] [全域会员通](https://open.taobao.com/API.htm?docId=46937&docType=2) | [github.com/bububa/opentaobao/api/alimember](https://pkg.go.dev/github.com/bububa/opentaobao/api/alimember) |
| [x] [交易猫API](https://open.taobao.com/API.htm?docId=46489&docType=2) | [github.com/bububa/opentaobao/api/jym](https://pkg.go.dev/github.com/bububa/opentaobao/api/jym) |
| [x] [影城自运营开放Api](https://open.taobao.com/API.htm?docId=46972&docType=2) | [github.com/bububa/opentaobao/api/filmtfavatar](https://pkg.go.dev/github.com/bububa/opentaobao/api/filmtfavatar) |
| [x] [商家应用](https://open.taobao.com/API.htm?docId=46995&docType=2) | [github.com/bububa/opentaobao/api/miniappcloud](https://pkg.go.dev/github.com/bububa/opentaobao/api/miniappcloud) |
| [x] [国际火车票API](https://open.taobao.com/API.htm?docId=47364&docType=2) | [github.com/bububa/opentaobao/api/rail](https://pkg.go.dev/github.com/bububa/opentaobao/api/rail) |
| [x] [天猫校园API](https://open.taobao.com/API.htm?docId=52079&docType=2) | [github.com/bububa/opentaobao/api/tmallcampus](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallcampus) |
| [x] [消息API](https://open.taobao.com/API.htm?docId=56166&docType=2) | [github.com/bububa/opentaobao/api/alimsg](https://pkg.go.dev/github.com/bububa/opentaobao/api/alimsg) |
| [x] [菜鸟发货工作台API](https://open.taobao.com/API.htm?docId=47727&docType=2) | [github.com/bububa/opentaobao/api/consignplatform](https://pkg.go.dev/github.com/bububa/opentaobao/api/consignplatform) |
| [x] [五道口-物流-自提API](https://open.taobao.com/API.htm?docId=47575&docType=2) | [github.com/bububa/opentaobao/api/wdklogistics](https://pkg.go.dev/github.com/bububa/opentaobao/api/wdklogistics) |
| [x] [企业订餐员工API](https://open.taobao.com/API.htm?docId=48679&docType=2) | [github.com/bububa/opentaobao/api/eleenterpriseemployee](https://pkg.go.dev/github.com/bububa/opentaobao/api/eleenterpriseemployee) |
| [x] [船票API](https://open.taobao.com/API.htm?docId=48055&docType=2) | [github.com/bububa/opentaobao/api/ship](https://pkg.go.dev/github.com/bububa/opentaobao/api/ship) |
| [x] [飞猪-综合交通api](https://open.taobao.com/API.htm?docId=53101&docType=2) | [github.com/bububa/opentaobao/api/alitripcar](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitripcar) |
| [x] [企业订餐店铺接口](https://open.taobao.com/API.htm?docId=49167&docType=2) | [github.com/bububa/opentaobao/api/eleenterpriserestaurant](https://pkg.go.dev/github.com/bububa/opentaobao/api/eleenterpriserestaurant) |
| [x] [新零售开放API](https://open.taobao.com/API.htm?docId=50091&docType=2) | [github.com/bububa/opentaobao/api/nropen](https://pkg.go.dev/github.com/bububa/opentaobao/api/nropen) |
| [x] [企业订餐购物车API](https://open.taobao.com/API.htm?docId=49011&docType=2) | [github.com/bububa/opentaobao/api/eleenterprisecartnew](https://pkg.go.dev/github.com/bububa/opentaobao/api/eleenterprisecartnew) |
| [x] [企业订餐优惠券API](https://open.taobao.com/API.htm?docId=49008&docType=2) | [github.com/bububa/opentaobao/api/eleenterprisecoupon](https://pkg.go.dev/github.com/bububa/opentaobao/api/eleenterprisecoupon) |
| [x] [阿里健康算法](https://open.taobao.com/API.htm?docId=48492&docType=2) | [github.com/bububa/opentaobao/api/alihealthalgo](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthalgo) |
| [x] [企业订餐订单API](https://open.taobao.com/API.htm?docId=49014&docType=2) | [github.com/bububa/opentaobao/api/eleenterpriseordernew](https://pkg.go.dev/github.com/bububa/opentaobao/api/eleenterpriseordernew) |
| [x] [云游戏API](https://open.taobao.com/API.htm?docId=55012&docType=2) | [github.com/bububa/opentaobao/api/cloudgame](https://pkg.go.dev/github.com/bububa/opentaobao/api/cloudgame) |
| [x] [国际虚拟API](https://open.taobao.com/API.htm?docId=48538&docType=2) | [github.com/bububa/opentaobao/api/globalvirtual](https://pkg.go.dev/github.com/bububa/opentaobao/api/globalvirtual) |
| [x] [交易开放](https://open.taobao.com/API.htm?docId=49615&docType=2) | [github.com/bububa/opentaobao/api/opentrade](https://pkg.go.dev/github.com/bububa/opentaobao/api/opentrade) |
| [x] [本地生活商户基础API](https://open.taobao.com/API.htm?docId=53173&docType=2) | [github.com/bububa/opentaobao/api/alscmerchant](https://pkg.go.dev/github.com/bububa/opentaobao/api/alscmerchant) |
| [x] [消息amp通道](https://open.taobao.com/API.htm?docId=50611&docType=2) | [github.com/bububa/opentaobao/api/msgamp](https://pkg.go.dev/github.com/bububa/opentaobao/api/msgamp) |
| [x] [口碑综合体API](https://open.taobao.com/API.htm?docId=49834&docType=2) | [github.com/bububa/opentaobao/api/koubeimall](https://pkg.go.dev/github.com/bububa/opentaobao/api/koubeimall) |
| [x] [IoT售后解决方案API](https://open.taobao.com/API.htm?docId=50245&docType=2) | [github.com/bububa/opentaobao/api/iotticket](https://pkg.go.dev/github.com/bububa/opentaobao/api/iotticket) |
| [x] [OpenMall-API](https://open.taobao.com/API.htm?docId=50836&docType=2) | [github.com/bububa/opentaobao/api/openmall](https://pkg.go.dev/github.com/bububa/opentaobao/api/openmall) |
| [x] [视觉开放API(viapi)](https://open.taobao.com/API.htm?docId=50726&docType=2) | [github.com/bububa/opentaobao/api/viapi](https://pkg.go.dev/github.com/bububa/opentaobao/api/viapi) |
| [x] [曲库开放平台歌曲API](https://open.taobao.com/API.htm?docId=54991&docType=2) | [github.com/bububa/opentaobao/api/xiamicontent](https://pkg.go.dev/github.com/bububa/opentaobao/api/xiamicontent) |
| [x] [阿里健康三方机构](https://open.taobao.com/API.htm?docId=51290&docType=2) | [github.com/bububa/opentaobao/api/alihealthmedical](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthmedical) |
| [x] [阿里健康-检测检验-预约](https://open.taobao.com/API.htm?docId=51444&docType=2) | [github.com/bububa/opentaobao/api/alihealthlab](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthlab) |
| [x] [天猫精灵供应链网关API](https://open.taobao.com/API.htm?docId=53192&docType=2) | [github.com/bububa/opentaobao/api/tmallgeniescp](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmallgeniescp) |
| [x] [曲库开放平台TraceAPI](https://open.taobao.com/API.htm?docId=51350&docType=2) | [github.com/bububa/opentaobao/api/xiamitrace](https://pkg.go.dev/github.com/bububa/opentaobao/api/xiamitrace) |
| [x] [飞猪发票](https://open.taobao.com/API.htm?docId=51546&docType=2) | [github.com/bububa/opentaobao/api/alitripreceipt](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitripreceipt) |
| [x] [互动开放API](https://open.taobao.com/API.htm?docId=53888&docType=2) | [github.com/bububa/opentaobao/api/jstinteractive](https://pkg.go.dev/github.com/bububa/opentaobao/api/jstinteractive) |
| [x] [同城零售全渠道](https://open.taobao.com/API.htm?docId=53988&docType=2) | [github.com/bububa/opentaobao/api/perfect](https://pkg.go.dev/github.com/bububa/opentaobao/api/perfect) |
| [x] [天猫好房工具API](https://open.taobao.com/API.htm?docId=56032&docType=2) | [github.com/bububa/opentaobao/api/alihouse](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihouse) |
| [x] [五棱镜任务API](https://open.taobao.com/API.htm?docId=52767&docType=2) | [github.com/bububa/opentaobao/api/pentraprism](https://pkg.go.dev/github.com/bububa/opentaobao/api/pentraprism) |
| [x] [闲鱼兼职](https://open.taobao.com/API.htm?docId=52622&docType=2) | [github.com/bububa/opentaobao/api/idleparttime](https://pkg.go.dev/github.com/bububa/opentaobao/api/idleparttime) |
| [x] [聚石塔隐私号](https://open.taobao.com/API.htm?docId=52636&docType=2) | [github.com/bububa/opentaobao/api/jstsecret](https://pkg.go.dev/github.com/bububa/opentaobao/api/jstsecret) |
| [x] [零售通商品API](https://open.taobao.com/API.htm?docId=47318&docType=2) | [github.com/bububa/opentaobao/api/lsticitem](https://pkg.go.dev/github.com/bububa/opentaobao/api/lsticitem) |
| [x] [零售通交易API](https://open.taobao.com/API.htm?docId=47800&docType=2) | [github.com/bububa/opentaobao/api/lsttrade](https://pkg.go.dev/github.com/bububa/opentaobao/api/lsttrade) |
| [x] [零售通门店API](https://open.taobao.com/API.htm?docId=43687&docType=2) | [github.com/bububa/opentaobao/api/lstbm](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstbm) |
| [x] [零售通履单API](https://open.taobao.com/API.htm?docId=53482&docType=2) | [github.com/bububa/opentaobao/api/lstlogistics2](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstlogistics2) |
| [x] [零售通仓储API](https://open.taobao.com/API.htm?docId=53981&docType=2) | [github.com/bububa/opentaobao/api/lstwarehouse](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstwarehouse) |
| [x] [零售通结算API](https://open.taobao.com/API.htm?docId=46608&docType=2) | [github.com/bububa/opentaobao/api/lstfundbill](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstfundbill) |
| [x] [零售通营销API](https://open.taobao.com/API.htm?docId=44185&docType=2) | [github.com/bububa/opentaobao/api/lstmarketing](https://pkg.go.dev/github.com/bububa/opentaobao/api/lstmarketing) |
| [x] [AE任务开放平台](https://open.taobao.com/API.htm?docId=53415&docType=2) | [github.com/bububa/opentaobao/api/aetask](https://pkg.go.dev/github.com/bububa/opentaobao/api/aetask) |
| [x] [本地生活内容API](https://open.taobao.com/API.htm?docId=54328&docType=2) | [github.com/bububa/opentaobao/api/kbalgo](https://pkg.go.dev/github.com/bububa/opentaobao/api/kbalgo) |
| [x] [天猫新品平台API](https://open.taobao.com/API.htm?docId=54824&docType=2) | [github.com/bububa/opentaobao/api/tmalltrend](https://pkg.go.dev/github.com/bububa/opentaobao/api/tmalltrend) |
| [x] [国际化中台服务域保险](https://open.taobao.com/API.htm?docId=53335&docType=2) | [github.com/bububa/opentaobao/api/middleclaims](https://pkg.go.dev/github.com/bububa/opentaobao/api/middleclaims) |
| [x] [阿里健康-健康证](https://open.taobao.com/API.htm?docId=54140&docType=2) | [github.com/bububa/opentaobao/api/alihealthcert](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthcert) |
| [x] [隐私保护API](https://open.taobao.com/API.htm?docId=54714&docType=2) | [github.com/bububa/opentaobao/api/topoaid](https://pkg.go.dev/github.com/bububa/opentaobao/api/topoaid) |
| [x] [ICBU-DropShipping](https://open.taobao.com/API.htm?docId=54908&docType=2) | [github.com/bububa/opentaobao/api/icbudropshipping](https://pkg.go.dev/github.com/bububa/opentaobao/api/icbudropshipping) |
| [x] [曲库开放平台艺人API](https://open.taobao.com/API.htm?docId=55125&docType=2) | [github.com/bububa/opentaobao/api/xiamiatrist](https://pkg.go.dev/github.com/bububa/opentaobao/api/xiamiatrist) |
| [x] [周期购API](https://open.taobao.com/API.htm?docId=55500&docType=2) | [github.com/bububa/opentaobao/api/zqs](https://pkg.go.dev/github.com/bububa/opentaobao/api/zqs) |
| [x] [闲鱼已验货](https://open.taobao.com/API.htm?docId=49487&docType=2) | [github.com/bububa/opentaobao/api/idleisv](https://pkg.go.dev/github.com/bububa/opentaobao/api/idleisv) |
| [x] [阿信-交易API](https://open.taobao.com/API.htm?docId=55695&docType=2) | [github.com/bububa/opentaobao/api/axintrade](https://pkg.go.dev/github.com/bububa/opentaobao/api/axintrade) |
| [x] [飞猪商业化API](https://open.taobao.com/API.htm?docId=55665&docType=2) | [github.com/bububa/opentaobao/api/alitripbp](https://pkg.go.dev/github.com/bububa/opentaobao/api/alitripbp) |
| [x] [阿信-基础数据](https://open.taobao.com/API.htm?docId=55968&docType=2) | [github.com/bububa/opentaobao/api/axindata](https://pkg.go.dev/github.com/bububa/opentaobao/api/axindata) |
| [x] [阿里健康公益线API](https://open.taobao.com/API.htm?docId=56009&docType=2) | [github.com/bububa/opentaobao/api/alihealthpw](https://pkg.go.dev/github.com/bububa/opentaobao/api/alihealthpw) |
| [x] [海南离岛对外API](https://open.taobao.com/API.htm?docId=56257&docType=2) | [github.com/bububa/opentaobao/api/dutyfree](https://pkg.go.dev/github.com/bububa/opentaobao/api/dutyfree) |

