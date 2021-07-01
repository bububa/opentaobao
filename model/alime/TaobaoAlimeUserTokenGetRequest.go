package alime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户免登录令牌 API请求
taobao.alime.user.token.get

根据第三账号信息获取用户的免登录令牌
*/
type TaobaoAlimeUserTokenGetAPIRequest struct {
    model.Params
    // 用户在第三方账号中的唯一id
    _foreignId   string
    // 用户昵称
    _nick   string
    // 小蜜分配给第三方账号的来源
    _source   int64
    // 用户在小蜜账号中的唯一id
    _id   int64
    // 令牌的过期时间(时间为秒)，最大为3600
    _expires   int64
    // 路由id, 一般为用户id，用于异地容灾
    _routing   int64
}

// 初始化TaobaoAlimeUserTokenGetAPIRequest对象
func NewTaobaoAlimeUserTokenGetRequest() *TaobaoAlimeUserTokenGetAPIRequest{
    return &TaobaoAlimeUserTokenGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlimeUserTokenGetAPIRequest) GetApiMethodName() string {
    return "taobao.alime.user.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlimeUserTokenGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetForeignId(_foreignId string) error {
    r._foreignId = _foreignId
    r.Set("foreign_id", _foreignId)
    return nil
}

// ForeignId Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetForeignId() string {
    return r._foreignId
}
// Nick Setter
// 用户昵称
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetNick() string {
    return r._nick
}
// Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetSource() int64 {
    return r._source
}
// Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetId() int64 {
    return r._id
}
// Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetExpires(_expires int64) error {
    r._expires = _expires
    r.Set("expires", _expires)
    return nil
}

// Expires Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetExpires() int64 {
    return r._expires
}
// Routing Setter
// 路由id, 一般为用户id，用于异地容灾
func (r *TaobaoAlimeUserTokenGetAPIRequest) SetRouting(_routing int64) error {
    r._routing = _routing
    r.Set("routing", _routing)
    return nil
}

// Routing Getter
func (r TaobaoAlimeUserTokenGetAPIRequest) GetRouting() int64 {
    return r._routing
}
