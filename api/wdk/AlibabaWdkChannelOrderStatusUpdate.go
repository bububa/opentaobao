package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
订单状态变更 
alibaba.wdk.channel.order.status.update

订单状态变更
*/
func AlibabaWdkChannelOrderStatusUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderStatusUpdateRequest, session string) (*wdk.AlibabaWdkChannelOrderStatusUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelOrderStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
