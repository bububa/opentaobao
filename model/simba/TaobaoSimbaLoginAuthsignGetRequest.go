package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆权限签名 APIRequest
taobao.simba.login.authsign.get

获取登陆权限签名
*/
type TaobaoSimbaLoginAuthsignGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaLoginAuthsignGetRequest() *TaobaoSimbaLoginAuthsignGetRequest{
    return &TaobaoSimbaLoginAuthsignGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaLoginAuthsignGetRequest) GetApiMethodName() string {
    return "taobao.simba.login.authsign.get"
}

func (r TaobaoSimbaLoginAuthsignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaLoginAuthsignGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaLoginAuthsignGetRequest) GetNick() string {
    return r.nick
}

