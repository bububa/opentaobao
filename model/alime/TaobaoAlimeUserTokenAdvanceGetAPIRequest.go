package alime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalimeusertokenadvancegetAPIRequest 获取用户免登录令牌v2 API请求
// taobao.alime.user.token.advance.get
//
// 根据第三账号信息获取用户的免登录令牌
type TaobaoalimeusertokenadvancegetAPIRequest struct {
	model.Params
	// 用户在第三方账号中的唯一id
	_foreignId string
	// 用户昵称
	_nick string
	// 路由id, 一般为用户id，用于异地容灾
	_routing int64
	// 用户类型，0为普通用户，1为访客用户
	_type int64
	// 小蜜分配给第三方账号的来源
	_source int64
	// 用户在小蜜账号中的唯一id
	_id int64
	// 令牌的过期时间(时间为秒)，最大为3600
	_expires int64
}

// NewTaobaoalimeusertokenadvancegetRequest 初始化TaobaoalimeusertokenadvancegetAPIRequest对象
func NewTaobaoalimeusertokenadvancegetRequest() *TaobaoalimeusertokenadvancegetAPIRequest {
	return &TaobaoalimeusertokenadvancegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetApiMethodName() string {
	return "taobao.alime.user.token.advance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetForeignId is ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetForeignId(_foreignId string) error {
	r._foreignId = _foreignId
	r.Set("foreign_id", _foreignId)
	return nil
}

// GetForeignId ForeignId Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetForeignId() string {
	return r._foreignId
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetNick() string {
	return r._nick
}

// SetRouting is Routing Setter
// 路由id, 一般为用户id，用于异地容灾
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetRouting(_routing int64) error {
	r._routing = _routing
	r.Set("routing", _routing)
	return nil
}

// GetRouting Routing Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetRouting() int64 {
	return r._routing
}

// SetType is Type Setter
// 用户类型，0为普通用户，1为访客用户
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetType() int64 {
	return r._type
}

// SetSource is Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetSource() int64 {
	return r._source
}

// SetId is Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetId() int64 {
	return r._id
}

// SetExpires is Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoalimeusertokenadvancegetAPIRequest) SetExpires(_expires int64) error {
	r._expires = _expires
	r.Set("expires", _expires)
	return nil
}

// GetExpires Expires Getter
func (r TaobaoalimeusertokenadvancegetAPIRequest) GetExpires() int64 {
	return r._expires
}
