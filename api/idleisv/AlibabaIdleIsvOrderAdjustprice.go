package idleisv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleisv"
)

/* 
闲鱼服务商订单价格修改接口 
alibaba.idle.isv.order.adjustprice

闲鱼用户通过授权的服务商修改订单价格和邮费
*/
func AlibabaIdleIsvOrderAdjustprice(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderAdjustpriceRequest, session string) (*idleisv.AlibabaIdleIsvOrderAdjustpriceAPIResponse, error) {
    var resp idleisv.AlibabaIdleIsvOrderAdjustpriceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
