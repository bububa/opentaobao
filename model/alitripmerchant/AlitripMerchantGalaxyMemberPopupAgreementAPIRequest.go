package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberpopupagreementAPIRequest 小程序唤起协议弹窗 API请求
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
type AlitripmerchantgalaxymemberpopupagreementAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 微信用户标识
	_code string
}

// NewAlitripmerchantgalaxymemberpopupagreementRequest 初始化AlitripmerchantgalaxymemberpopupagreementAPIRequest对象
func NewAlitripmerchantgalaxymemberpopupagreementRequest() *AlitripmerchantgalaxymemberpopupagreementAPIRequest {
	return &AlitripmerchantgalaxymemberpopupagreementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberpopupagreementAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.popup.agreement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberpopupagreementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberpopupagreementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxymemberpopupagreementAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberpopupagreementAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCode is Code Setter
// 微信用户标识
func (r *AlitripmerchantgalaxymemberpopupagreementAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripmerchantgalaxymemberpopupagreementAPIRequest) GetCode() string {
	return r._code
}
