package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
消息发送记录查询 
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。
*/
func TaobaoTmcMsgSendrecord(clt *core.SDKClient, req *tmc.TaobaoTmcMsgSendrecordRequest, session string) (*tmc.TaobaoTmcMsgSendrecordResponse, error) {
    var resp tmc.TaobaoTmcMsgSendrecordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
