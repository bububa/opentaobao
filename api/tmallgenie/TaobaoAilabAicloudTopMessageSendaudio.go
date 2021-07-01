package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
发送语音留言 
taobao.ailab.aicloud.top.message.sendaudio

将语音的二进制byte[]通过TOP接口发送保存
*/
func TaobaoAilabAicloudTopMessageSendaudio(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessageSendaudioAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMessageSendaudioAPIResponse, error) {
    var resp tmallgenie.TaobaoAilabAicloudTopMessageSendaudioAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
