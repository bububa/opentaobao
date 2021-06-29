package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
创建揽收物流 
alibaba.idle.rent.order.logistics.deliver

创建揽收物流
商家去物流公司创建物流订单
*/
func AlibabaIdleRentOrderLogisticsDeliver(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderLogisticsDeliverRequest, session string) (*idle.AlibabaIdleRentOrderLogisticsDeliverAPIResponse, error) {
    var resp idle.AlibabaIdleRentOrderLogisticsDeliverAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
