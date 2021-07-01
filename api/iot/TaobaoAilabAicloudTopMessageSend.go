package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
发送留言 
taobao.ailab.aicloud.top.message.send

供准入的外部用户实现发送留言功能，APP端发送，设备端读取
*/
func TaobaoAilabAicloudTopMessageSend(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageSendAPIRequest, session string) (*iot.TaobaoAilabAicloudTopMessageSendAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopMessageSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
