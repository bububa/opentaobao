package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageAuthCheckAPIRequest 校验用户授权关系 API请求
// taobao.top.package.auth.check
//
// 校验用户授权关系
type TaobaoTopPackageAuthCheckAPIRequest struct {
	model.Params
	// 权限校验入参
	_authScopeCheckRequest *AuthScopeCheckRequest
}

// NewTaobaoTopPackageAuthCheckRequest 初始化TaobaoTopPackageAuthCheckAPIRequest对象
func NewTaobaoTopPackageAuthCheckRequest() *TaobaoTopPackageAuthCheckAPIRequest {
	return &TaobaoTopPackageAuthCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopPackageAuthCheckAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.auth.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopPackageAuthCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopPackageAuthCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthScopeCheckRequest is AuthScopeCheckRequest Setter
// 权限校验入参
func (r *TaobaoTopPackageAuthCheckAPIRequest) SetAuthScopeCheckRequest(_authScopeCheckRequest *AuthScopeCheckRequest) error {
	r._authScopeCheckRequest = _authScopeCheckRequest
	r.Set("auth_scope_check_request", _authScopeCheckRequest)
	return nil
}

// GetAuthScopeCheckRequest AuthScopeCheckRequest Getter
func (r TaobaoTopPackageAuthCheckAPIRequest) GetAuthScopeCheckRequest() *AuthScopeCheckRequest {
	return r._authScopeCheckRequest
}
