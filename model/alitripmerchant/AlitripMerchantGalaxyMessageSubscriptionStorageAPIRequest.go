package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest 存储模版ID API请求
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
type AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest struct {
	model.Params
	// 模版ID
	_templateIds string
	// 租户ID
	_tenantKey string
	// token
	_token string
}

// NewAlitripMerchantGalaxyMessageSubscriptionStorageRequest 初始化AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest对象
func NewAlitripMerchantGalaxyMessageSubscriptionStorageRequest() *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest {
	return &AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) Reset() {
	r._templateIds = ""
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.message.subscription.storage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateIds is TemplateIds Setter
// 模版ID
func (r *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) SetTemplateIds(_templateIds string) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// GetTemplateIds TemplateIds Getter
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetTemplateIds() string {
	return r._templateIds
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMessageSubscriptionStorageRequest()
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionStorageRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest
func GetAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest() *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest {
	return poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest.Get().(*AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest 将 AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest(v *AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest.Put(v)
}
