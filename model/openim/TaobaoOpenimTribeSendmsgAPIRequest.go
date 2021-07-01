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
type TaobaoOpenimTribeSendmsgAPIRequest struct {
    model.Params
    // 群消息发送者，只有该群的成员才可以发送群消息
    _user   *User
    // 群id
    _tribeId   int64
    // 发送群消息
    _msg   *TribeMsg
}

// 初始化TaobaoOpenimTribeSendmsgAPIRequest对象
func NewTaobaoOpenimTribeSendmsgRequest() *TaobaoOpenimTribeSendmsgAPIRequest{
    return &TaobaoOpenimTribeSendmsgAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.sendmsg"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 群消息发送者，只有该群的成员才可以发送群消息
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetUser(_user *User) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetUser() *User {
    return r._user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetTribeId() int64 {
    return r._tribeId
}
// Msg Setter
// 发送群消息
func (r *TaobaoOpenimTribeSendmsgAPIRequest) SetMsg(_msg *TribeMsg) error {
    r._msg = _msg
    r.Set("msg", _msg)
    return nil
}

// Msg Getter
func (r TaobaoOpenimTribeSendmsgAPIRequest) GetMsg() *TribeMsg {
    return r._msg
}
