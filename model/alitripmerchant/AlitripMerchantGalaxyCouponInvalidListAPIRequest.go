package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCouponInvalidListAPIRequest 用户失效优惠券列表 API请求
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
type AlitripMerchantGalaxyCouponInvalidListAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyCouponInvalidListRequest 初始化AlitripMerchantGalaxyCouponInvalidListAPIRequest对象
func NewAlitripMerchantGalaxyCouponInvalidListRequest() *AlitripMerchantGalaxyCouponInvalidListAPIRequest {
	return &AlitripMerchantGalaxyCouponInvalidListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyCouponInvalidListAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.coupon.invalid.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyCouponInvalidListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyCouponInvalidListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyCouponInvalidListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyCouponInvalidListRequest()
	},
}

// GetAlitripMerchantGalaxyCouponInvalidListRequest 从 sync.Pool 获取 AlitripMerchantGalaxyCouponInvalidListAPIRequest
func GetAlitripMerchantGalaxyCouponInvalidListAPIRequest() *AlitripMerchantGalaxyCouponInvalidListAPIRequest {
	return poolAlitripMerchantGalaxyCouponInvalidListAPIRequest.Get().(*AlitripMerchantGalaxyCouponInvalidListAPIRequest)
}

// ReleaseAlitripMerchantGalaxyCouponInvalidListAPIRequest 将 AlitripMerchantGalaxyCouponInvalidListAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyCouponInvalidListAPIRequest(v *AlitripMerchantGalaxyCouponInvalidListAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyCouponInvalidListAPIRequest.Put(v)
}
