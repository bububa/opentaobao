package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
分页拉取订单数据 
alibaba.wdk.trade.order.balance.bill.query

提供接口供外部调用，分页拉取订单数据
*/
func AlibabaWdkTradeOrderBalanceBillQuery(clt *core.SDKClient, req *wdk.AlibabaWdkTradeOrderBalanceBillQueryRequest, session string) (*wdk.AlibabaWdkTradeOrderBalanceBillQueryResponse, error) {
    var resp wdk.AlibabaWdkTradeOrderBalanceBillQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
