package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitygoodsqueryAPIRequest 营销抽奖-用户奖品查询 API请求
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
type AlitripmerchantgalaxyactivitygoodsqueryAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录标识
	_token string
}

// NewAlitripmerchantgalaxyactivitygoodsqueryRequest 初始化AlitripmerchantgalaxyactivitygoodsqueryAPIRequest对象
func NewAlitripmerchantgalaxyactivitygoodsqueryRequest() *AlitripmerchantgalaxyactivitygoodsqueryAPIRequest {
	return &AlitripmerchantgalaxyactivitygoodsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.goods.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitygoodsqueryAPIRequest) GetToken() string {
	return r._token
}
