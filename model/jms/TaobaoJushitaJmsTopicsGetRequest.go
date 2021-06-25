package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户nick获取开通的消息列表 APIRequest
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
type TaobaoJushitaJmsTopicsGetRequest struct {
    model.Params

    // 卖家nick
    nick   string 

}

func NewTaobaoJushitaJmsTopicsGetRequest() *TaobaoJushitaJmsTopicsGetRequest{
    return &TaobaoJushitaJmsTopicsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsTopicsGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.topics.get"
}

func (r TaobaoJushitaJmsTopicsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsTopicsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoJushitaJmsTopicsGetRequest) GetNick() string {
    return r.nick
}

