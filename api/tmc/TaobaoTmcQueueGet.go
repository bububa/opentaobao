package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
获取消息队列积压情况 
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况
*/
func TaobaoTmcQueueGet(clt *core.SDKClient, req *tmc.TaobaoTmcQueueGetAPIRequest, session string) (*tmc.TaobaoTmcQueueGetAPIResponse, error) {
    var resp tmc.TaobaoTmcQueueGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
