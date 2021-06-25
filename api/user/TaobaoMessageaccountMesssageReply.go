package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
消息号下行回复接口 
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复
*/
func TaobaoMessageaccountMesssageReply(clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageReplyRequest, session string) (*user.TaobaoMessageaccountMesssageReplyResponse, error) {
    var resp user.TaobaoMessageaccountMesssageReplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
