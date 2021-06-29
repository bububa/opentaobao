package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消用户的消息服务 API请求
taobao.tmc.user.cancel

取消用户的消息服务
*/
type TaobaoTmcUserCancelRequest struct {
    model.Params
    // 用户昵称
    _nick   string
    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    _userPlatform   string
}

// 初始化TaobaoTmcUserCancelRequest对象
func NewTaobaoTmcUserCancelRequest() *TaobaoTmcUserCancelRequest{
    return &TaobaoTmcUserCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserCancelRequest) GetApiMethodName() string {
    return "taobao.tmc.user.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户昵称
func (r *TaobaoTmcUserCancelRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoTmcUserCancelRequest) GetNick() string {
    return r._nick
}
// UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserCancelRequest) SetUserPlatform(_userPlatform string) error {
    r._userPlatform = _userPlatform
    r.Set("user_platform", _userPlatform)
    return nil
}

// UserPlatform Getter
func (r TaobaoTmcUserCancelRequest) GetUserPlatform() string {
    return r._userPlatform
}
