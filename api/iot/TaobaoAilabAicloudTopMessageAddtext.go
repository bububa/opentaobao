package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
精灵代说 
taobao.ailab.aicloud.top.message.addtext

精灵代说
*/
func TaobaoAilabAicloudTopMessageAddtext(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageAddtextRequest, session string) (*iot.TaobaoAilabAicloudTopMessageAddtextAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopMessageAddtextAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
