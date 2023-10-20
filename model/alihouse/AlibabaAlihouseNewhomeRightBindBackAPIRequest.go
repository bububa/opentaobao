package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeRightBindBackAPIRequest) Reset() {
	r._userInfo = nil
	r.Params.ToZero()
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

var poolAlibabaAlihouseNewhomeRightBindBackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeRightBindBackRequest()
	},
}

// GetAlibabaAlihouseNewhomeRightBindBackRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeRightBindBackAPIRequest
func GetAlibabaAlihouseNewhomeRightBindBackAPIRequest() *AlibabaAlihouseNewhomeRightBindBackAPIRequest {
	return poolAlibabaAlihouseNewhomeRightBindBackAPIRequest.Get().(*AlibabaAlihouseNewhomeRightBindBackAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeRightBindBackAPIRequest 将 AlibabaAlihouseNewhomeRightBindBackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRightBindBackAPIRequest(v *AlibabaAlihouseNewhomeRightBindBackAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRightBindBackAPIRequest.Put(v)
}
