package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest 查询用户是否有模版ID权限 API请求
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
type AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// token
	_token string
}

// NewAlitripMerchantGalaxyMessageSubscriptionQueryRequest 初始化AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest对象
func NewAlitripMerchantGalaxyMessageSubscriptionQueryRequest() *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest {
	return &AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.message.subscription.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMessageSubscriptionQueryRequest()
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest
func GetAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest() *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest {
	return poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest.Get().(*AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest 将 AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest(v *AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest.Put(v)
}
