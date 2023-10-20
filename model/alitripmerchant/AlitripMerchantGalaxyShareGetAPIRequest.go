package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxysharegetAPIRequest 星河-获取小程序分享文案和图片 API请求
// alitrip.merchant.galaxy.share.get
//
// 获取 雅高微信小程序分享素材文案和图片。
type AlitripmerchantgalaxysharegetAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
}

// NewAlitripmerchantgalaxysharegetRequest 初始化AlitripmerchantgalaxysharegetAPIRequest对象
func NewAlitripmerchantgalaxysharegetRequest() *AlitripmerchantgalaxysharegetAPIRequest {
	return &AlitripmerchantgalaxysharegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxysharegetAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.share.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxysharegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxysharegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxysharegetAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxysharegetAPIRequest) GetTenantKey() string {
	return r._tenantKey
}
