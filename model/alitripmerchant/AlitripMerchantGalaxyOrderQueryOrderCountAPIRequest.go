package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest 查询各种状态订单的总数 API请求
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
type AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyOrderQueryOrderCountRequest 初始化AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest对象
func NewAlitripMerchantGalaxyOrderQueryOrderCountRequest() *AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest {
	return &AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query.order.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderQueryOrderCountAPIRequest) GetToken() string {
	return r._token
}
