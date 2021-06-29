package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户nick获取开通的消息列表 API请求
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
type TaobaoJushitaJmsTopicsGetRequest struct {
    model.Params
    // 卖家nick
    nick   string
}

// 初始化TaobaoJushitaJmsTopicsGetRequest对象
func NewTaobaoJushitaJmsTopicsGetRequest() *TaobaoJushitaJmsTopicsGetRequest{
    return &TaobaoJushitaJmsTopicsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsTopicsGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.topics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsTopicsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 卖家nick
func (r *TaobaoJushitaJmsTopicsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoJushitaJmsTopicsGetRequest) GetNick() string {
    return r.nick
}
