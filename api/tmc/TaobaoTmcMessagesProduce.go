package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
批量发送消息 
taobao.tmc.messages.produce

批量发送消息
*/
func TaobaoTmcMessagesProduce(clt *core.SDKClient, req *tmc.TaobaoTmcMessagesProduceAPIRequest, session string) (*tmc.TaobaoTmcMessagesProduceAPIResponse, error) {
    var resp tmc.TaobaoTmcMessagesProduceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
