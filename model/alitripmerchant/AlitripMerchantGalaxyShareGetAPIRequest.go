package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyShareGetAPIRequest 星河-获取小程序分享文案和图片 API请求
// alitrip.merchant.galaxy.share.get
//
// 获取 雅高微信小程序分享素材文案和图片。
type AlitripMerchantGalaxyShareGetAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
}

// NewAlitripMerchantGalaxyShareGetRequest 初始化AlitripMerchantGalaxyShareGetAPIRequest对象
func NewAlitripMerchantGalaxyShareGetRequest() *AlitripMerchantGalaxyShareGetAPIRequest {
	return &AlitripMerchantGalaxyShareGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyShareGetAPIRequest) Reset() {
	r._tenantKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.share.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyShareGetAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

var poolAlitripMerchantGalaxyShareGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyShareGetRequest()
	},
}

// GetAlitripMerchantGalaxyShareGetRequest 从 sync.Pool 获取 AlitripMerchantGalaxyShareGetAPIRequest
func GetAlitripMerchantGalaxyShareGetAPIRequest() *AlitripMerchantGalaxyShareGetAPIRequest {
	return poolAlitripMerchantGalaxyShareGetAPIRequest.Get().(*AlitripMerchantGalaxyShareGetAPIRequest)
}

// ReleaseAlitripMerchantGalaxyShareGetAPIRequest 将 AlitripMerchantGalaxyShareGetAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyShareGetAPIRequest(v *AlitripMerchantGalaxyShareGetAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyShareGetAPIRequest.Put(v)
}
