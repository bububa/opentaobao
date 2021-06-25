package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行回复接口 APIRequest
taobao.miniapp.messsage.reply

外部 isv 调用该进口来进行轻店铺消息的回复
*/
type TaobaoMiniappMesssageReplyRequest struct {
    model.Params

    // 入参
    param   *ReplyMessageDto 

}

func NewTaobaoMiniappMesssageReplyRequest() *TaobaoMiniappMesssageReplyRequest{
    return &TaobaoMiniappMesssageReplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappMesssageReplyRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.reply"
}

func (r TaobaoMiniappMesssageReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappMesssageReplyRequest) SetParam(param *ReplyMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoMiniappMesssageReplyRequest) GetParam() *ReplyMessageDto {
    return r.param
}

