package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityMarketingPopupAPIRequest 营销弹屏 API请求
// alitrip.merchant.galaxy.activity.marketing.popup
//
// 星河=活动营销弹屏
type AlitripMerchantGalaxyActivityMarketingPopupAPIRequest struct {
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

// NewAlitripMerchantGalaxyActivityMarketingPopupRequest 初始化AlitripMerchantGalaxyActivityMarketingPopupAPIRequest对象
func NewAlitripMerchantGalaxyActivityMarketingPopupRequest() *AlitripMerchantGalaxyActivityMarketingPopupAPIRequest {
	return &AlitripMerchantGalaxyActivityMarketingPopupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.marketing.popup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetToken() string {
	return r._token
}

// SetCode is Code Setter
// 用户code
func (r *AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetCode() string {
	return r._code
}

// SetVersion is Version Setter
// 版本，适配灰度发布
func (r *AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlitripMerchantGalaxyActivityMarketingPopupAPIRequest) GetVersion() string {
	return r._version
}
