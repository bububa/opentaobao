package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaloginauthsigngetAPIRequest 获取登陆权限签名 API请求
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
type TaobaosimbaloginauthsigngetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// NewTaobaosimbaloginauthsigngetRequest 初始化TaobaosimbaloginauthsigngetAPIRequest对象
func NewTaobaosimbaloginauthsigngetRequest() *TaobaosimbaloginauthsigngetAPIRequest {
	return &TaobaosimbaloginauthsigngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaloginauthsigngetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.login.authsign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaloginauthsigngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaloginauthsigngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaloginauthsigngetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaloginauthsigngetAPIRequest) GetNick() string {
	return r._nick
}
