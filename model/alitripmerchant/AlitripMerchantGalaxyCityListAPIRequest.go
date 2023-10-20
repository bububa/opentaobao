package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycitylistAPIRequest 星河-酒店城市列表展示 API请求
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
type AlitripmerchantgalaxycitylistAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 0国内 1国外
	_domestic int64
}

// NewAlitripmerchantgalaxycitylistRequest 初始化AlitripmerchantgalaxycitylistAPIRequest对象
func NewAlitripmerchantgalaxycitylistRequest() *AlitripmerchantgalaxycitylistAPIRequest {
	return &AlitripmerchantgalaxycitylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycitylistAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycitylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycitylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxycitylistAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycitylistAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDomestic is Domestic Setter
// 0国内 1国外
func (r *AlitripmerchantgalaxycitylistAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r AlitripmerchantgalaxycitylistAPIRequest) GetDomestic() int64 {
	return r._domestic
}
