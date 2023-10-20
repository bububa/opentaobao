package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatusermodifyphoneAPIRequest DFC-ID用户换绑手机号 API请求
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
type AlitripmerchantgalaxywechatusermodifyphoneAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 微信向程序登录参数
	_loginParam *LoginParam
}

// NewAlitripmerchantgalaxywechatusermodifyphoneRequest 初始化AlitripmerchantgalaxywechatusermodifyphoneAPIRequest对象
func NewAlitripmerchantgalaxywechatusermodifyphoneRequest() *AlitripmerchantgalaxywechatusermodifyphoneAPIRequest {
	return &AlitripmerchantgalaxywechatusermodifyphoneAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.user.modify.phone"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetToken() string {
	return r._token
}

// SetLoginParam is LoginParam Setter
// 微信向程序登录参数
func (r *AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripmerchantgalaxywechatusermodifyphoneAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}
