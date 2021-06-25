package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户是否同步消息 APIRequest
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询
*/
type TaobaoJushitaJmsUserGetRequest struct {
    model.Params

    // 需要查询的用户名
    userNick   string 

}

func NewTaobaoJushitaJmsUserGetRequest() *TaobaoJushitaJmsUserGetRequest{
    return &TaobaoJushitaJmsUserGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsUserGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.get"
}

func (r TaobaoJushitaJmsUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsUserGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoJushitaJmsUserGetRequest) GetUserNick() string {
    return r.userNick
}

