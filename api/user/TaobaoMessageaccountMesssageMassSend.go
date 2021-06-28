package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
消息号开放-消息群发 
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发
*/
func TaobaoMessageaccountMesssageMassSend(clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageMassSendRequest, session string) (*user.TaobaoMessageaccountMesssageMassSendAPIResponse, error) {
    var resp user.TaobaoMessageaccountMesssageMassSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
