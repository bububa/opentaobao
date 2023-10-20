package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest 查询用户是否有模版ID权限 API请求
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
type AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// token
	_token string
}

// NewAlitripmerchantgalaxymessagesubscriptionqueryRequest 初始化AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest对象
func NewAlitripmerchantgalaxymessagesubscriptionqueryRequest() *AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest {
	return &AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.message.subscription.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest) GetToken() string {
	return r._token
}
