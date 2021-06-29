package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
闲鱼无需物流发货 
alibaba.idle.order.dummy.send

适用于电子卡券等虚拟商品不需要物流的商品发货。
*/
func AlibabaIdleOrderDummySend(clt *core.SDKClient, req *idle.AlibabaIdleOrderDummySendRequest, session string) (*idle.AlibabaIdleOrderDummySendAPIResponse, error) {
    var resp idle.AlibabaIdleOrderDummySendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
