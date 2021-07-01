package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
帮卖订单履约 
alibaba.idle.consignment.order.perform

帮卖订单履约，回收商同步订单信息，驱动交易流转
*/
func AlibabaIdleConsignmentOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentOrderPerformAPIRequest, session string) (*idle.AlibabaIdleConsignmentOrderPerformAPIResponse, error) {
    var resp idle.AlibabaIdleConsignmentOrderPerformAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
