package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderFillAPIRequest 填单页接口 API请求
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
type AlitripMerchantGalaxyOrderFillAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 填单页入参
	_fillOrderParam *FillOrderParam
}

// NewAlitripMerchantGalaxyOrderFillRequest 初始化AlitripMerchantGalaxyOrderFillAPIRequest对象
func NewAlitripMerchantGalaxyOrderFillRequest() *AlitripMerchantGalaxyOrderFillAPIRequest {
	return &AlitripMerchantGalaxyOrderFillAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyOrderFillAPIRequest) Reset() {
	r._tenantKey = ""
	r._fillOrderParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderFillAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.fill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderFillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOrderFillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderFillAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderFillAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetFillOrderParam is FillOrderParam Setter
// 填单页入参
func (r *AlitripMerchantGalaxyOrderFillAPIRequest) SetFillOrderParam(_fillOrderParam *FillOrderParam) error {
	r._fillOrderParam = _fillOrderParam
	r.Set("fill_order_param", _fillOrderParam)
	return nil
}

// GetFillOrderParam FillOrderParam Getter
func (r AlitripMerchantGalaxyOrderFillAPIRequest) GetFillOrderParam() *FillOrderParam {
	return r._fillOrderParam
}

var poolAlitripMerchantGalaxyOrderFillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyOrderFillRequest()
	},
}

// GetAlitripMerchantGalaxyOrderFillRequest 从 sync.Pool 获取 AlitripMerchantGalaxyOrderFillAPIRequest
func GetAlitripMerchantGalaxyOrderFillAPIRequest() *AlitripMerchantGalaxyOrderFillAPIRequest {
	return poolAlitripMerchantGalaxyOrderFillAPIRequest.Get().(*AlitripMerchantGalaxyOrderFillAPIRequest)
}

// ReleaseAlitripMerchantGalaxyOrderFillAPIRequest 将 AlitripMerchantGalaxyOrderFillAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderFillAPIRequest(v *AlitripMerchantGalaxyOrderFillAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderFillAPIRequest.Put(v)
}
