package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
小程序活动短信发送 
taobao.jst.miniapp.crowd.message.send

小程序活动短信发送
*/
func TaobaoJstMiniappCrowdMessageSend(clt *core.SDKClient, req *jst.TaobaoJstMiniappCrowdMessageSendRequest, session string) (*jst.TaobaoJstMiniappCrowdMessageSendResponse, error) {
    var resp jst.TaobaoJstMiniappCrowdMessageSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
