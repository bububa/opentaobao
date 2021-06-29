package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵消息中心单播推送消息接口 
taobao.ailab.aicloud.top.message.push.unicast

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
*/
func TaobaoAilabAicloudTopMessagePushUnicast(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessagePushUnicastRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMessagePushUnicastAPIResponse, error) {
    var resp tmallgenie.TaobaoAilabAicloudTopMessagePushUnicastAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
