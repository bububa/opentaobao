package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRightBindBackAPIRequest 权限回流 API请求
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
type AlibabaAlihouseNewhomeRightBindBackAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserPermissionInfoDto
}

// NewAlibabaAlihouseNewhomeRightBindBackRequest 初始化AlibabaAlihouseNewhomeRightBindBackAPIRequest对象
func NewAlibabaAlihouseNewhomeRightBindBackRequest() *AlibabaAlihouseNewhomeRightBindBackAPIRequest {
	return &AlibabaAlihouseNewhomeRightBindBackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRightBindBackAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.right.bind.back"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRightBindBackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeRightBindBackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaAlihouseNewhomeRightBindBackAPIRequest) SetUserInfo(_userInfo *UserPermissionInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaAlihouseNewhomeRightBindBackAPIRequest) GetUserInfo() *UserPermissionInfoDto {
	return r._userInfo
}
