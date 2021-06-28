package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
创建订单 
alibaba.wdk.channel.order.create

外部商家创建订单
*/
func AlibabaWdkChannelOrderCreate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderCreateRequest, session string) (*wdk.AlibabaWdkChannelOrderCreateAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
