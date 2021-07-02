package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyProviderMemberQueryAPIRequest 提供会员查询接口 API请求
// alitrip.merchant.galaxy.provider.member.query
//
// 星河产品=提供会查询服务
type AlitripMerchantGalaxyProviderMemberQueryAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 查询参数
	_queryMemberParam *QueryMemberParam
}

// NewAlitripMerchantGalaxyProviderMemberQueryRequest 初始化AlitripMerchantGalaxyProviderMemberQueryAPIRequest对象
func NewAlitripMerchantGalaxyProviderMemberQueryRequest() *AlitripMerchantGalaxyProviderMemberQueryAPIRequest {
	return &AlitripMerchantGalaxyProviderMemberQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.provider.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is QueryMemberParam Setter
// 查询参数
func (r *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) SetQueryMemberParam(_queryMemberParam *QueryMemberParam) error {
	r._queryMemberParam = _queryMemberParam
	r.Set("query_member_param", _queryMemberParam)
	return nil
}

// Get QueryMemberParam Getter
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetQueryMemberParam() *QueryMemberParam {
	return r._queryMemberParam
}
