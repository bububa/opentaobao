package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝用户头像查询 APIRequest
taobao.user.avatar.get

根据混淆nick查询用户头像
*/
type TaobaoUserAvatarGetRequest struct {
    model.Params

    // 混淆nick
    nick   string 

}

func NewTaobaoUserAvatarGetRequest() *TaobaoUserAvatarGetRequest{
    return &TaobaoUserAvatarGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUserAvatarGetRequest) GetApiMethodName() string {
    return "taobao.user.avatar.get"
}

func (r TaobaoUserAvatarGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUserAvatarGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoUserAvatarGetRequest) GetNick() string {
    return r.nick
}

