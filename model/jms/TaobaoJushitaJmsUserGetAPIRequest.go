package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajmsusergetAPIRequest 查询某个用户是否同步消息 API请求
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
type TaobaojushitajmsusergetAPIRequest struct {
	model.Params
	// 需要查询的用户名
	_userNick string
}

// NewTaobaojushitajmsusergetRequest 初始化TaobaojushitajmsusergetAPIRequest对象
func NewTaobaojushitajmsusergetRequest() *TaobaojushitajmsusergetAPIRequest {
	return &TaobaojushitajmsusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajmsusergetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajmsusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajmsusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 需要查询的用户名
func (r *TaobaojushitajmsusergetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaojushitajmsusergetAPIRequest) GetUserNick() string {
	return r._userNick
}
