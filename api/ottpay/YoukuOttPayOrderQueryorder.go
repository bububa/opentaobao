package ottpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ottpay"
)

/* 
查询订单 
youku.ott.pay.order.queryorder

通过订单号查询订单信息
*/
func YoukuOttPayOrderQueryorder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryorderAPIRequest, session string) (*ottpay.YoukuOttPayOrderQueryorderAPIResponse, error) {
    var resp ottpay.YoukuOttPayOrderQueryorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
