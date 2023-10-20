package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderqueryordercountAPIRequest 查询各种状态订单的总数 API请求
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
type AlitripmerchantgalaxyorderqueryordercountAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripmerchantgalaxyorderqueryordercountRequest 初始化AlitripmerchantgalaxyorderqueryordercountAPIRequest对象
func NewAlitripmerchantgalaxyorderqueryordercountRequest() *AlitripmerchantgalaxyorderqueryordercountAPIRequest {
	return &AlitripmerchantgalaxyorderqueryordercountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyorderqueryordercountAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query.order.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyorderqueryordercountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyorderqueryordercountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyorderqueryordercountAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyorderqueryordercountAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyorderqueryordercountAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyorderqueryordercountAPIRequest) GetToken() string {
	return r._token
}
