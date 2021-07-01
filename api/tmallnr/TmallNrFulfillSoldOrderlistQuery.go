package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
零售商获取品牌商的特定订单列表 
tmall.nr.fulfill.sold.orderlist.query

零售商获取品牌商的特定订单列表，后端服务有零售商和品牌商的绑定关系，存在开关控制；同时后端存在定时送业务等特殊业务的校验，非同城配送业务不能返回，返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
*/
func TmallNrFulfillSoldOrderlistQuery(clt *core.SDKClient, req *tmallnr.TmallNrFulfillSoldOrderlistQueryAPIRequest, session string) (*tmallnr.TmallNrFulfillSoldOrderlistQueryAPIResponse, error) {
    var resp tmallnr.TmallNrFulfillSoldOrderlistQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
