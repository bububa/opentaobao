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
    routing   int64
    // 用户类型，0为普通用户，1为访客用户
    type   int64
    // 用户在第三方账号中的唯一id
    foreignId   string
    // 用户昵称
    nick   string
    // 小蜜分配给第三方账号的来源
    source   int64
    // 用户在小蜜账号中的唯一id
    id   int64
    // 令牌的过期时间(时间为秒)，最大为3600
    expires   int64
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
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetRouting(routing int64) error {
    r.routing = routing
    r.Set("routing", routing)
    return nil
}

// Routing Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetRouting() int64 {
    return r.routing
}
// Type Setter
// 用户类型，0为普通用户，1为访客用户
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetType() int64 {
    return r.type
}
// ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetForeignId(foreignId string) error {
    r.foreignId = foreignId
    r.Set("foreign_id", foreignId)
    return nil
}

// ForeignId Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetForeignId() string {
    return r.foreignId
}
// Nick Setter
// 用户昵称
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetNick() string {
    return r.nick
}
// Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetSource() int64 {
    return r.source
}
// Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetId() int64 {
    return r.id
}
// Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoAlimeUserTokenAdvanceGetRequest) SetExpires(expires int64) error {
    r.expires = expires
    r.Set("expires", expires)
    return nil
}

// Expires Getter
func (r TaobaoAlimeUserTokenAdvanceGetRequest) GetExpires() int64 {
    return r.expires
}
