package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomerightbindbackAPIRequest 权限回流 API请求
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
type AlibabaalihousenewhomerightbindbackAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserPermissionInfoDto
}

// NewAlibabaalihousenewhomerightbindbackRequest 初始化AlibabaalihousenewhomerightbindbackAPIRequest对象
func NewAlibabaalihousenewhomerightbindbackRequest() *AlibabaalihousenewhomerightbindbackAPIRequest {
	return &AlibabaalihousenewhomerightbindbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomerightbindbackAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.right.bind.back"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomerightbindbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomerightbindbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaalihousenewhomerightbindbackAPIRequest) SetUserInfo(_userInfo *UserPermissionInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaalihousenewhomerightbindbackAPIRequest) GetUserInfo() *UserPermissionInfoDto {
	return r._userInfo
}
