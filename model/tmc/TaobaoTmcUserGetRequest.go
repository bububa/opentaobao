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
    fields   string
    // 用户昵称
    nick   string
    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    userPlatform   string
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
func (r *TaobaoTmcUserGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTmcUserGetRequest) GetFields() string {
    return r.fields
}
// Nick Setter
// 用户昵称
func (r *TaobaoTmcUserGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoTmcUserGetRequest) GetNick() string {
    return r.nick
}
// UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserGetRequest) SetUserPlatform(userPlatform string) error {
    r.userPlatform = userPlatform
    r.Set("user_platform", userPlatform)
    return nil
}

// UserPlatform Getter
func (r TaobaoTmcUserGetRequest) GetUserPlatform() string {
    return r.userPlatform
}
