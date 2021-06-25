package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号下行回复接口 APIRequest
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复
*/
type TaobaoMessageaccountMesssageReplyRequest struct {
    model.Params

    // 入参
    param0   *ReplyMessageDto 

}

func NewTaobaoMessageaccountMesssageReplyRequest() *TaobaoMessageaccountMesssageReplyRequest{
    return &TaobaoMessageaccountMesssageReplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMessageaccountMesssageReplyRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.reply"
}

func (r TaobaoMessageaccountMesssageReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMessageaccountMesssageReplyRequest) SetParam0(param0 *ReplyMessageDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoMessageaccountMesssageReplyRequest) GetParam0() *ReplyMessageDto {
    return r.param0
}

