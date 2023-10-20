package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymembertokenAPIRequest 星河-校验token API请求
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
type AlitripmerchantgalaxymembertokenAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 访问携带toekn
	_token string
}

// NewAlitripmerchantgalaxymembertokenRequest 初始化AlitripmerchantgalaxymembertokenAPIRequest对象
func NewAlitripmerchantgalaxymembertokenRequest() *AlitripmerchantgalaxymembertokenAPIRequest {
	return &AlitripmerchantgalaxymembertokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymembertokenAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.token"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymembertokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymembertokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxymembertokenAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymembertokenAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 访问携带toekn
func (r *AlitripmerchantgalaxymembertokenAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymembertokenAPIRequest) GetToken() string {
	return r._token
}
