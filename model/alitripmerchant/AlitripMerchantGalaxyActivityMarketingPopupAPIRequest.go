package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitymarketingpopupAPIRequest 营销弹屏 API请求
// alitrip.merchant.galaxy.activity.marketing.popup
//
// 星河=活动营销弹屏
type AlitripmerchantgalaxyactivitymarketingpopupAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 用户code
	_code string
	// 版本，适配灰度发布
	_version string
}

// NewAlitripmerchantgalaxyactivitymarketingpopupRequest 初始化AlitripmerchantgalaxyactivitymarketingpopupAPIRequest对象
func NewAlitripmerchantgalaxyactivitymarketingpopupRequest() *AlitripmerchantgalaxyactivitymarketingpopupAPIRequest {
	return &AlitripmerchantgalaxyactivitymarketingpopupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.marketing.popup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetToken() string {
	return r._token
}

// SetCode is Code Setter
// 用户code
func (r *AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetCode() string {
	return r._code
}

// SetVersion is Version Setter
// 版本，适配灰度发布
func (r *AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlitripmerchantgalaxyactivitymarketingpopupAPIRequest) GetVersion() string {
	return r._version
}
