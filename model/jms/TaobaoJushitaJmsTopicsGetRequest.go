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
type TaobaoJushitaJmsTopicsGetAPIRequest struct {
    model.Params
    // 卖家nick
    _nick   string
}

// 初始化TaobaoJushitaJmsTopicsGetAPIRequest对象
func NewTaobaoJushitaJmsTopicsGetRequest() *TaobaoJushitaJmsTopicsGetAPIRequest{
    return &TaobaoJushitaJmsTopicsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.topics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 卖家nick
func (r *TaobaoJushitaJmsTopicsGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoJushitaJmsTopicsGetAPIRequest) GetNick() string {
    return r._nick
}
