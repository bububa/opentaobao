package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyBrandSearchAPIRequest 星河-品牌搜索 API请求
// alitrip.merchant.galaxy.brand.search
//
// 星河服务=获取雅高品牌信息
type AlitripMerchantGalaxyBrandSearchAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
}

// NewAlitripMerchantGalaxyBrandSearchRequest 初始化AlitripMerchantGalaxyBrandSearchAPIRequest对象
func NewAlitripMerchantGalaxyBrandSearchRequest() *AlitripMerchantGalaxyBrandSearchAPIRequest {
	return &AlitripMerchantGalaxyBrandSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyBrandSearchAPIRequest) Reset() {
	r._tenantKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.brand.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyBrandSearchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

var poolAlitripMerchantGalaxyBrandSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyBrandSearchRequest()
	},
}

// GetAlitripMerchantGalaxyBrandSearchRequest 从 sync.Pool 获取 AlitripMerchantGalaxyBrandSearchAPIRequest
func GetAlitripMerchantGalaxyBrandSearchAPIRequest() *AlitripMerchantGalaxyBrandSearchAPIRequest {
	return poolAlitripMerchantGalaxyBrandSearchAPIRequest.Get().(*AlitripMerchantGalaxyBrandSearchAPIRequest)
}

// ReleaseAlitripMerchantGalaxyBrandSearchAPIRequest 将 AlitripMerchantGalaxyBrandSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyBrandSearchAPIRequest(v *AlitripMerchantGalaxyBrandSearchAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyBrandSearchAPIRequest.Put(v)
}
