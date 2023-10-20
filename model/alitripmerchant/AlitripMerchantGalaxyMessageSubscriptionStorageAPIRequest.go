package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest 存储模版ID API请求
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
type AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest struct {
	model.Params
	// 模版ID
	_templateIds string
	// 租户ID
	_tenantKey string
	// token
	_token string
}

// NewAlitripmerchantgalaxymessagesubscriptionstorageRequest 初始化AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest对象
func NewAlitripmerchantgalaxymessagesubscriptionstorageRequest() *AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest {
	return &AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.message.subscription.storage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateIds is TemplateIds Setter
// 模版ID
func (r *AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) SetTemplateIds(_templateIds string) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// GetTemplateIds TemplateIds Getter
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetTemplateIds() string {
	return r._templateIds
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest) GetToken() string {
	return r._token
}
