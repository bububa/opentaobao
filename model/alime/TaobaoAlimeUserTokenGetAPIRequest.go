package alime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalimeusertokengetAPIRequest 获取用户免登录令牌 API请求
// taobao.alime.user.token.get
//
// 根据第三账号信息获取用户的免登录令牌
type TaobaoalimeusertokengetAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 用户在第三方账号中的唯一id
	_foreignId string
	// 路由id, 一般为用户id，用于异地容灾
	_routing int64
	// 用户在小蜜账号中的唯一id
	_id int64
	// 小蜜分配给第三方账号的来源
	_source int64
	// 令牌的过期时间(时间为秒)，最大为3600
	_expires int64
}

// NewTaobaoalimeusertokengetRequest 初始化TaobaoalimeusertokengetAPIRequest对象
func NewTaobaoalimeusertokengetRequest() *TaobaoalimeusertokengetAPIRequest {
	return &TaobaoalimeusertokengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalimeusertokengetAPIRequest) GetApiMethodName() string {
	return "taobao.alime.user.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalimeusertokengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalimeusertokengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoalimeusertokengetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoalimeusertokengetAPIRequest) GetNick() string {
	return r._nick
}

// SetForeignId is ForeignId Setter
// 用户在第三方账号中的唯一id
func (r *TaobaoalimeusertokengetAPIRequest) SetForeignId(_foreignId string) error {
	r._foreignId = _foreignId
	r.Set("foreign_id", _foreignId)
	return nil
}

// GetForeignId ForeignId Getter
func (r TaobaoalimeusertokengetAPIRequest) GetForeignId() string {
	return r._foreignId
}

// SetRouting is Routing Setter
// 路由id, 一般为用户id，用于异地容灾
func (r *TaobaoalimeusertokengetAPIRequest) SetRouting(_routing int64) error {
	r._routing = _routing
	r.Set("routing", _routing)
	return nil
}

// GetRouting Routing Getter
func (r TaobaoalimeusertokengetAPIRequest) GetRouting() int64 {
	return r._routing
}

// SetId is Id Setter
// 用户在小蜜账号中的唯一id
func (r *TaobaoalimeusertokengetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoalimeusertokengetAPIRequest) GetId() int64 {
	return r._id
}

// SetSource is Source Setter
// 小蜜分配给第三方账号的来源
func (r *TaobaoalimeusertokengetAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoalimeusertokengetAPIRequest) GetSource() int64 {
	return r._source
}

// SetExpires is Expires Setter
// 令牌的过期时间(时间为秒)，最大为3600
func (r *TaobaoalimeusertokengetAPIRequest) SetExpires(_expires int64) error {
	r._expires = _expires
	r.Set("expires", _expires)
	return nil
}

// GetExpires Expires Getter
func (r TaobaoalimeusertokengetAPIRequest) GetExpires() int64 {
	return r._expires
}
