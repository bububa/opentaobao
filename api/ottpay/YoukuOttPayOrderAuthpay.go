package ottpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ottpay"
)

/* 
委托代扣服务 
youku.ott.pay.order.authpay

应用中心sdk连续包月委托代扣服务
*/
func YoukuOttPayOrderAuthpay(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderAuthpayRequest, session string) (*ottpay.YoukuOttPayOrderAuthpayAPIResponse, error) {
    var resp ottpay.YoukuOttPayOrderAuthpayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
