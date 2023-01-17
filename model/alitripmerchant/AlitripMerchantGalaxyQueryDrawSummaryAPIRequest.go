package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyQueryDrawSummaryAPIRequest 星河-抽奖活动概要列表查询 API请求
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
type AlitripMerchantGalaxyQueryDrawSummaryAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
}

// NewAlitripMerchantGalaxyQueryDrawSummaryRequest 初始化AlitripMerchantGalaxyQueryDrawSummaryAPIRequest对象
func NewAlitripMerchantGalaxyQueryDrawSummaryRequest() *AlitripMerchantGalaxyQueryDrawSummaryAPIRequest {
	return &AlitripMerchantGalaxyQueryDrawSummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.query.draw.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyQueryDrawSummaryAPIRequest) GetToken() string {
	return r._token
}
