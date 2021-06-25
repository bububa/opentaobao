package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS消息同步用户 APIRequest
taobao.jushita.jms.user.delete

删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
*/
type TaobaoJushitaJmsUserDeleteRequest struct {
    model.Params

    // 需要停止同步消息的用户nick
    userNick   string 

}

func NewTaobaoJushitaJmsUserDeleteRequest() *TaobaoJushitaJmsUserDeleteRequest{
    return &TaobaoJushitaJmsUserDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsUserDeleteRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.delete"
}

func (r TaobaoJushitaJmsUserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsUserDeleteRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoJushitaJmsUserDeleteRequest) GetUserNick() string {
    return r.userNick
}

