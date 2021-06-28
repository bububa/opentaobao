package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
批量发送奇门事件 
taobao.qimen.events.produce

批量发送消息
*/
func TaobaoQimenEventsProduce(clt *core.SDKClient, req *util.TaobaoQimenEventsProduceRequest, session string) (*util.TaobaoQimenEventsProduceAPIResponse, error) {
    var resp util.TaobaoQimenEventsProduceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
