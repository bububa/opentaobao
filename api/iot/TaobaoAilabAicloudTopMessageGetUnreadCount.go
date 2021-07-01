package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取未读的消息数量 
taobao.ailab.aicloud.top.message.get.unread.count

开放获取未读留言数量的接口
*/
func TaobaoAilabAicloudTopMessageGetUnreadCount(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest, session string) (*iot.TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
