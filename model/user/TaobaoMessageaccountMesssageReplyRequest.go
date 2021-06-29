package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号下行回复接口 API请求
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复
*/
type TaobaoMessageaccountMesssageReplyRequest struct {
    model.Params
    // 入参
    param0   *ReplyMessageDto
}

// 初始化TaobaoMessageaccountMesssageReplyRequest对象
func NewTaobaoMessageaccountMesssageReplyRequest() *TaobaoMessageaccountMesssageReplyRequest{
    return &TaobaoMessageaccountMesssageReplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageReplyRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.reply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TaobaoMessageaccountMesssageReplyRequest) SetParam0(param0 *ReplyMessageDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoMessageaccountMesssageReplyRequest) GetParam0() *ReplyMessageDto {
    return r.param0
}
