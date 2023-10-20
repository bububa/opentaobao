package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotoppackageauthcheckAPIRequest 校验用户授权关系 API请求
// taobao.top.package.auth.check
//
// 校验用户授权关系
type TaobaotoppackageauthcheckAPIRequest struct {
	model.Params
	// 权限校验入参
	_authScopeCheckRequest *AuthScopeCheckRequest
}

// NewTaobaotoppackageauthcheckRequest 初始化TaobaotoppackageauthcheckAPIRequest对象
func NewTaobaotoppackageauthcheckRequest() *TaobaotoppackageauthcheckAPIRequest {
	return &TaobaotoppackageauthcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotoppackageauthcheckAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.auth.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotoppackageauthcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotoppackageauthcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthScopeCheckRequest is AuthScopeCheckRequest Setter
// 权限校验入参
func (r *TaobaotoppackageauthcheckAPIRequest) SetAuthScopeCheckRequest(_authScopeCheckRequest *AuthScopeCheckRequest) error {
	r._authScopeCheckRequest = _authScopeCheckRequest
	r.Set("auth_scope_check_request", _authScopeCheckRequest)
	return nil
}

// GetAuthScopeCheckRequest AuthScopeCheckRequest Getter
func (r TaobaotoppackageauthcheckAPIRequest) GetAuthScopeCheckRequest() *AuthScopeCheckRequest {
	return r._authScopeCheckRequest
}
