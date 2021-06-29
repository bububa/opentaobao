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
type TaobaoAlimeUserTokenGetRequest struct {
    model.Params
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
    // 路由id, 一般为用户id，用于异地容灾
    routing   int64
}

// 初始化TaobaoAlimeUserTokenGetRequest对象
func NewTaobaoAlimeUserTokenGetRequest() *TaobaoAlimeUserTokenGetRequest{
    return &TaobaoAlimeUserTokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlimeUserTokenGetRequest) GetApiMethodName() string {
    return "taobao.alime.user.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlimeUserTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoAlimeUserTokenGetRequest) SetForeignId(foreignId string) error {
    r.foreignId = foreignId
    r.Set("foreign_id", foreignId)
    return nil
}

// ForeignId Getter
func (r TaobaoAlimeUserTokenGetRequest) GetForeignId() string {
    return r.foreignId
}
// Nick Setter
// 用户昵称
func (r *TaobaoAlimeUserTokenGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoAlimeUserTokenGetRequest) GetNick() string {
    return r.nick
}
// Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoAlimeUserTokenGetRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoAlimeUserTokenGetRequest) GetSource() int64 {
    return r.source
}
// Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoAlimeUserTokenGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoAlimeUserTokenGetRequest) GetId() int64 {
    return r.id
}
// Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoAlimeUserTokenGetRequest) SetExpires(expires int64) error {
    r.expires = expires
    r.Set("expires", expires)
    return nil
}

// Expires Getter
func (r TaobaoAlimeUserTokenGetRequest) GetExpires() int64 {
    return r.expires
}
// Routing Setter
// 路由id, 一般为用户id，用于异地容灾
func (r *TaobaoAlimeUserTokenGetRequest) SetRouting(routing int64) error {
    r.routing = routing
    r.Set("routing", routing)
    return nil
}

// Routing Getter
func (r TaobaoAlimeUserTokenGetRequest) GetRouting() int64 {
    return r.routing
}
