package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
确认发货 
alibaba.idle.rent.order.senditem

确认发货
*/
func AlibabaIdleRentOrderSenditem(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderSenditemRequest, session string) (*idle.AlibabaIdleRentOrderSenditemAPIResponse, error) {
    var resp idle.AlibabaIdleRentOrderSenditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
