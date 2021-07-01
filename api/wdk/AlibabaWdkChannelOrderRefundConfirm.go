package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
退款确认 
alibaba.wdk.channel.order.refund.confirm

退款确认
*/
func AlibabaWdkChannelOrderRefundConfirm(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderRefundConfirmAPIRequest, session string) (*wdk.AlibabaWdkChannelOrderRefundConfirmAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelOrderRefundConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
