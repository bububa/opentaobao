package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
零售云小程序消息推送 
alibaba.lsy.miniapp.msg.push

零售云小程序消息推送，推送消息至零售云（喵零等）
*/
func AlibabaLsyMiniappMsgPush(clt *core.SDKClient, req *tmc.AlibabaLsyMiniappMsgPushAPIRequest, session string) (*tmc.AlibabaLsyMiniappMsgPushAPIResponse, error) {
    var resp tmc.AlibabaLsyMiniappMsgPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
