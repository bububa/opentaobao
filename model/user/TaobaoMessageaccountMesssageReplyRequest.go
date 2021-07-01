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
type TaobaoMessageaccountMesssageReplyAPIRequest struct {
    model.Params
    // 入参
    _param0   *ReplyMessageDTO
}

// 初始化TaobaoMessageaccountMesssageReplyAPIRequest对象
func NewTaobaoMessageaccountMesssageReplyRequest() *TaobaoMessageaccountMesssageReplyAPIRequest{
    return &TaobaoMessageaccountMesssageReplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.reply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TaobaoMessageaccountMesssageReplyAPIRequest) SetParam0(_param0 *ReplyMessageDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoMessageaccountMesssageReplyAPIRequest) GetParam0() *ReplyMessageDTO {
    return r._param0
}
