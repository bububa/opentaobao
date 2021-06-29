package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户已开通消息 API请求
taobao.tmc.user.get

查询指定用户开通的消息通道和组
*/
type TaobaoTmcUserGetRequest struct {
    model.Params
    // 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
    _fields   string
    // 用户昵称
    _nick   string
    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    _userPlatform   string
}

// 初始化TaobaoTmcUserGetRequest对象
func NewTaobaoTmcUserGetRequest() *TaobaoTmcUserGetRequest{
    return &TaobaoTmcUserGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserGetRequest) GetApiMethodName() string {
    return "taobao.tmc.user.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
func (r *TaobaoTmcUserGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTmcUserGetRequest) GetFields() string {
    return r._fields
}
// Nick Setter
// 用户昵称
func (r *TaobaoTmcUserGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoTmcUserGetRequest) GetNick() string {
    return r._nick
}
// UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserGetRequest) SetUserPlatform(_userPlatform string) error {
    r._userPlatform = _userPlatform
    r.Set("user_platform", _userPlatform)
    return nil
}

// UserPlatform Getter
func (r TaobaoTmcUserGetRequest) GetUserPlatform() string {
    return r._userPlatform
}
