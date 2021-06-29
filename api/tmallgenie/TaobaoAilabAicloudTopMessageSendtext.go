package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
故事机发送文本留言 
taobao.ailab.aicloud.top.message.sendtext

故事机文本留言
*/
func TaobaoAilabAicloudTopMessageSendtext(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessageSendtextRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMessageSendtextAPIResponse, error) {
    var resp tmallgenie.TaobaoAilabAicloudTopMessageSendtextAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
