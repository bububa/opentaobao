package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝用户头像查询 API请求
taobao.user.avatar.get

根据混淆nick查询用户头像
*/
type TaobaoUserAvatarGetRequest struct {
    model.Params
    // 混淆nick
    nick   string
}

// 初始化TaobaoUserAvatarGetRequest对象
func NewTaobaoUserAvatarGetRequest() *TaobaoUserAvatarGetRequest{
    return &TaobaoUserAvatarGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUserAvatarGetRequest) GetApiMethodName() string {
    return "taobao.user.avatar.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUserAvatarGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 混淆nick
func (r *TaobaoUserAvatarGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoUserAvatarGetRequest) GetNick() string {
    return r.nick
}
