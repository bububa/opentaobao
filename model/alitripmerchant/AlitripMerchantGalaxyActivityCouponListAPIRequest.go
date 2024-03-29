package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityCouponListAPIRequest 用户领券中心列表 API请求
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
type AlitripMerchantGalaxyActivityCouponListAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户Token
	_token string
}

// NewAlitripMerchantGalaxyActivityCouponListRequest 初始化AlitripMerchantGalaxyActivityCouponListAPIRequest对象
func NewAlitripMerchantGalaxyActivityCouponListRequest() *AlitripMerchantGalaxyActivityCouponListAPIRequest {
	return &AlitripMerchantGalaxyActivityCouponListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityCouponListAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.coupon.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripMerchantGalaxyActivityCouponListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户Token
func (r *AlitripMerchantGalaxyActivityCouponListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyActivityCouponListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityCouponListRequest()
	},
}

// GetAlitripMerchantGalaxyActivityCouponListRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityCouponListAPIRequest
func GetAlitripMerchantGalaxyActivityCouponListAPIRequest() *AlitripMerchantGalaxyActivityCouponListAPIRequest {
	return poolAlitripMerchantGalaxyActivityCouponListAPIRequest.Get().(*AlitripMerchantGalaxyActivityCouponListAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityCouponListAPIRequest 将 AlitripMerchantGalaxyActivityCouponListAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityCouponListAPIRequest(v *AlitripMerchantGalaxyActivityCouponListAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityCouponListAPIRequest.Put(v)
}
