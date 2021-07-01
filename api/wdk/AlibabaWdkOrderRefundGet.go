package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口订单退款按ID查询 
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情
*/
func AlibabaWdkOrderRefundGet(clt *core.SDKClient, req *wdk.AlibabaWdkOrderRefundGetAPIRequest, session string) (*wdk.AlibabaWdkOrderRefundGetAPIResponse, error) {
    var resp wdk.AlibabaWdkOrderRefundGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
