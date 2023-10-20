package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyprovidermemberqueryAPIRequest 提供会员查询接口 API请求
// alitrip.merchant.galaxy.provider.member.query
//
// 星河产品=提供会查询服务
type AlitripmerchantgalaxyprovidermemberqueryAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 查询参数
	_queryMemberParam *QueryMemberParam
}

// NewAlitripmerchantgalaxyprovidermemberqueryRequest 初始化AlitripmerchantgalaxyprovidermemberqueryAPIRequest对象
func NewAlitripmerchantgalaxyprovidermemberqueryRequest() *AlitripmerchantgalaxyprovidermemberqueryAPIRequest {
	return &AlitripmerchantgalaxyprovidermemberqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyprovidermemberqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.provider.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyprovidermemberqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyprovidermemberqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyprovidermemberqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyprovidermemberqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetQueryMemberParam is QueryMemberParam Setter
// 查询参数
func (r *AlitripmerchantgalaxyprovidermemberqueryAPIRequest) SetQueryMemberParam(_queryMemberParam *QueryMemberParam) error {
	r._queryMemberParam = _queryMemberParam
	r.Set("query_member_param", _queryMemberParam)
	return nil
}

// GetQueryMemberParam QueryMemberParam Getter
func (r AlitripmerchantgalaxyprovidermemberqueryAPIRequest) GetQueryMemberParam() *QueryMemberParam {
	return r._queryMemberParam
}
