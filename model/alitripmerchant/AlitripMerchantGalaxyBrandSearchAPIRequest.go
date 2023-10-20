package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxybrandsearchAPIRequest 星河-品牌搜索 API请求
// alitrip.merchant.galaxy.brand.search
//
// 星河服务=获取雅高品牌信息
type AlitripmerchantgalaxybrandsearchAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
}

// NewAlitripmerchantgalaxybrandsearchRequest 初始化AlitripmerchantgalaxybrandsearchAPIRequest对象
func NewAlitripmerchantgalaxybrandsearchRequest() *AlitripmerchantgalaxybrandsearchAPIRequest {
	return &AlitripmerchantgalaxybrandsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxybrandsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.brand.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxybrandsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxybrandsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户信息
func (r *AlitripmerchantgalaxybrandsearchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxybrandsearchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}
