package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUserDeleteAPIRequest 删除数据推送用户 API请求
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteAPIRequest struct {
	model.Params
	// 要删除用户的昵称
	_nick string
	// 需要删除的用户编号
	_userId int64
}

// NewTaobaoJushitaJdpUserDeleteRequest 初始化TaobaoJushitaJdpUserDeleteAPIRequest对象
func NewTaobaoJushitaJdpUserDeleteRequest() *TaobaoJushitaJdpUserDeleteAPIRequest {
	return &TaobaoJushitaJdpUserDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 要删除用户的昵称
func (r *TaobaoJushitaJdpUserDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetNick() string {
	return r._nick
}

// Set is UserId Setter
// 需要删除的用户编号
func (r *TaobaoJushitaJdpUserDeleteAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetUserId() int64 {
	return r._userId
}
