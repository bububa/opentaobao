package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
翱象商家自有渠道 订单创建 
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建
*/
func AlibabaTclsAelophyMerchantChannelOrderCreate(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest, session string) (*wdk.AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse, error) {
    var resp wdk.AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
