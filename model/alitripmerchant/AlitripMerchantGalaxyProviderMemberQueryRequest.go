package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供会员查询接口 API请求
alitrip.merchant.galaxy.provider.member.query

星河产品=提供会查询服务
*/
type AlitripMerchantGalaxyProviderMemberQueryRequest struct {
    model.Params
    // 租户id
    tenantKey   string
    // 查询参数
    queryMemberParam   *QueryMemberParam
}

// 初始化AlitripMerchantGalaxyProviderMemberQueryRequest对象
func NewAlitripMerchantGalaxyProviderMemberQueryRequest() *AlitripMerchantGalaxyProviderMemberQueryRequest{
    return &AlitripMerchantGalaxyProviderMemberQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyProviderMemberQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.provider.member.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyProviderMemberQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyProviderMemberQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyProviderMemberQueryRequest) GetTenantKey() string {
    return r.tenantKey
}
// QueryMemberParam Setter
// 查询参数
func (r *AlitripMerchantGalaxyProviderMemberQueryRequest) SetQueryMemberParam(queryMemberParam *QueryMemberParam) error {
    r.queryMemberParam = queryMemberParam
    r.Set("query_member_param", queryMemberParam)
    return nil
}

// QueryMemberParam Getter
func (r AlitripMerchantGalaxyProviderMemberQueryRequest) GetQueryMemberParam() *QueryMemberParam {
    return r.queryMemberParam
}
