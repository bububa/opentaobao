package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
翱象商家自有渠道 逆向单完成 
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
func AlibabaTclsAelophyMerchantChannelRefundComplete(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelRefundCompleteRequest, session string) (*wdk.AlibabaTclsAelophyMerchantChannelRefundCompleteResponse, error) {
    var resp wdk.AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
