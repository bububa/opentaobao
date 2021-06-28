package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
发出奇门事件 
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。
*/
func TaobaoQimenEventProduce(clt *core.SDKClient, req *util.TaobaoQimenEventProduceRequest, session string) (*util.TaobaoQimenEventProduceAPIResponse, error) {
    var resp util.TaobaoQimenEventProduceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
