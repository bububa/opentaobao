package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送群消息 API请求
taobao.openim.tribe.sendmsg

发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
*/
type TaobaoOpenimTribeSendmsgRequest struct {
    model.Params
    // 群消息发送者，只有该群的成员才可以发送群消息
    user   *User
    // 群id
    tribeId   int64
    // 发送群消息
    msg   *TribeMsg
}

// 初始化TaobaoOpenimTribeSendmsgRequest对象
func NewTaobaoOpenimTribeSendmsgRequest() *TaobaoOpenimTribeSendmsgRequest{
    return &TaobaoOpenimTribeSendmsgRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSendmsgRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.sendmsg"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSendmsgRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 群消息发送者，只有该群的成员才可以发送群消息
func (r *TaobaoOpenimTribeSendmsgRequest) SetUser(user *User) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeSendmsgRequest) GetUser() *User {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeSendmsgRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeSendmsgRequest) GetTribeId() int64 {
    return r.tribeId
}
// Msg Setter
// 发送群消息
func (r *TaobaoOpenimTribeSendmsgRequest) SetMsg(msg *TribeMsg) error {
    r.msg = msg
    r.Set("msg", msg)
    return nil
}

// Msg Getter
func (r TaobaoOpenimTribeSendmsgRequest) GetMsg() *TribeMsg {
    return r.msg
}
