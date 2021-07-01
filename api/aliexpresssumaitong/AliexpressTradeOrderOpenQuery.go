package aliexpresssumaitong

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

/* 
Aliexpress开放平台订单查询 
aliexpress.trade.order.open.query

Aliexpress开放平台订单信息查询
*/
func AliexpressTradeOrderOpenQuery(clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTradeOrderOpenQueryAPIRequest, session string) (*aliexpresssumaitong.AliexpressTradeOrderOpenQueryAPIResponse, error) {
    var resp aliexpresssumaitong.AliexpressTradeOrderOpenQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
