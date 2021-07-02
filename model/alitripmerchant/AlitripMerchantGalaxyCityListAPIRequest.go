package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCityListAPIRequest 星河-酒店城市列表展示 API请求
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
type AlitripMerchantGalaxyCityListAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 0国内 1国外
	_domestic int64
}

// NewAlitripMerchantGalaxyCityListRequest 初始化AlitripMerchantGalaxyCityListAPIRequest对象
func NewAlitripMerchantGalaxyCityListRequest() *AlitripMerchantGalaxyCityListAPIRequest {
	return &AlitripMerchantGalaxyCityListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyCityListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is Domestic Setter
// 0国内 1国外
func (r *AlitripMerchantGalaxyCityListAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// Get Domestic Getter
func (r AlitripMerchantGalaxyCityListAPIRequest) GetDomestic() int64 {
	return r._domestic
}
