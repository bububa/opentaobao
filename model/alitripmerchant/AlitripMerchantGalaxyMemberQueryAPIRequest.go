package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberqueryAPIRequest 星河-获取登录用户的信息 API请求
// alitrip.merchant.galaxy.member.query
//
// 获取登录用户的信息
type AlitripmerchantgalaxymemberqueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// toekn
	_token string
}

// NewAlitripmerchantgalaxymemberqueryRequest 初始化AlitripmerchantgalaxymemberqueryAPIRequest对象
func NewAlitripmerchantgalaxymemberqueryRequest() *AlitripmerchantgalaxymemberqueryAPIRequest {
	return &AlitripmerchantgalaxymemberqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxymemberqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// toekn
func (r *AlitripmerchantgalaxymemberqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymemberqueryAPIRequest) GetToken() string {
	return r._token
}
