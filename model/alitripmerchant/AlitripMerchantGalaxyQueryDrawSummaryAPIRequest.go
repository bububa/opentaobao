package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyquerydrawsummaryAPIRequest 星河-抽奖活动概要列表查询 API请求
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
type AlitripmerchantgalaxyquerydrawsummaryAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
}

// NewAlitripmerchantgalaxyquerydrawsummaryRequest 初始化AlitripmerchantgalaxyquerydrawsummaryAPIRequest对象
func NewAlitripmerchantgalaxyquerydrawsummaryRequest() *AlitripmerchantgalaxyquerydrawsummaryAPIRequest {
	return &AlitripmerchantgalaxyquerydrawsummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyquerydrawsummaryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.query.draw.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyquerydrawsummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyquerydrawsummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxyquerydrawsummaryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyquerydrawsummaryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripmerchantgalaxyquerydrawsummaryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyquerydrawsummaryAPIRequest) GetToken() string {
	return r._token
}
