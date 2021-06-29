package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
alibaba查询订单支付结果 
alibaba.order.pay.result.query

alibaba查询订单支付结果
*/
func AlibabaOrderPayResultQuery(clt *core.SDKClient, req *icbudropshipping.AlibabaOrderPayResultQueryRequest, session string) (*icbudropshipping.AlibabaOrderPayResultQueryAPIResponse, error) {
    var resp icbudropshipping.AlibabaOrderPayResultQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
