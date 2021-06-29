package ottpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ottpay"
)

/* 
订单查询接口(cp订单号查询) 
youku.ott.pay.order.queryorderbycp

给商户服务端查询订单状态
*/
func YoukuOttPayOrderQueryorderbycp(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryorderbycpRequest, session string) (*ottpay.YoukuOttPayOrderQueryorderbycpAPIResponse, error) {
    var resp ottpay.YoukuOttPayOrderQueryorderbycpAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}