package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomerightunbindbackAPIRequest 权限回流-解绑 API请求
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
type AlibabaalihousenewhomerightunbindbackAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserPermissionInfoDto
}

// NewAlibabaalihousenewhomerightunbindbackRequest 初始化AlibabaalihousenewhomerightunbindbackAPIRequest对象
func NewAlibabaalihousenewhomerightunbindbackRequest() *AlibabaalihousenewhomerightunbindbackAPIRequest {
	return &AlibabaalihousenewhomerightunbindbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomerightunbindbackAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.right.unbind.back"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomerightunbindbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomerightunbindbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaalihousenewhomerightunbindbackAPIRequest) SetUserInfo(_userInfo *UserPermissionInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaalihousenewhomerightunbindbackAPIRequest) GetUserInfo() *UserPermissionInfoDto {
	return r._userInfo
}
