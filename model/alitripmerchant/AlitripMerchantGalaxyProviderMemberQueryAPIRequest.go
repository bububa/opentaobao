package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._queryMemberParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.provider.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetQueryMemberParam is QueryMemberParam Setter
// 查询参数
func (r *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) SetQueryMemberParam(_queryMemberParam *QueryMemberParam) error {
	r._queryMemberParam = _queryMemberParam
	r.Set("query_member_param", _queryMemberParam)
	return nil
}

// GetQueryMemberParam QueryMemberParam Getter
func (r AlitripMerchantGalaxyProviderMemberQueryAPIRequest) GetQueryMemberParam() *QueryMemberParam {
	return r._queryMemberParam
}

var poolAlitripMerchantGalaxyProviderMemberQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyProviderMemberQueryRequest()
	},
}

// GetAlitripMerchantGalaxyProviderMemberQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyProviderMemberQueryAPIRequest
func GetAlitripMerchantGalaxyProviderMemberQueryAPIRequest() *AlitripMerchantGalaxyProviderMemberQueryAPIRequest {
	return poolAlitripMerchantGalaxyProviderMemberQueryAPIRequest.Get().(*AlitripMerchantGalaxyProviderMemberQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyProviderMemberQueryAPIRequest 将 AlitripMerchantGalaxyProviderMemberQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyProviderMemberQueryAPIRequest(v *AlitripMerchantGalaxyProviderMemberQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyProviderMemberQueryAPIRequest.Put(v)
}
