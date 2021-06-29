package alime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户免登录令牌v2 API请求
taobao.alime.user.token.advance.get

根据第三账号信息获取用户的免登录令牌
*/
type TaobaoAlimeUserTokenAdvanceGetRequest struct {
    model.Params
    // 路由id, 一般为用户id，用于异地容灾
    _routing   int64
    // 用户类型，0为普通用户，1为访客用户
    _type   int64
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
}

// 初始化TaobaoAlimeUserTokenAdvanceGetRequest对象
func NewTaobaoAlimeUserTokenAdvanceGetRequest() *TaobaoAlimeUserTokenAdvanceGetRequest{
    return &TaobaoAlimeUserTokenAdvanceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetApiMethodName() string {
    return "taobao.alime.user.token.advance.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Routing Setter
// 路由id, 一般为用户id，用于异地容灾
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetRouting(_routing int64) error {
    r._routing = _routing
    r.Set("routing", _routing)
    return nil
}

// Routing Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetRouting() int64 {
    return r._routing
}
// Type Setter
// 用户类型，0为普通用户，1为访客用户
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetType() int64 {
    return r._type
}
// ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetForeignId(_foreignId string) error {
    r._foreignId = _foreignId
    r.Set("foreign_id", _foreignId)
    return nil
}

// ForeignId Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetForeignId() string {
    return r._foreignId
}
// Nick Setter
// 用户昵称
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetNick() string {
    return r._nick
}
// Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetSource() int64 {
    return r._source
}
// Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetId() int64 {
    return r._id
}
// Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetExpires(_expires int64) error {
    r._expires = _expires
    r.Set("expires", _expires)
    return nil
}

// Expires Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetExpires() int64 {
    return r._expires
}
