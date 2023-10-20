package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest 查询用户拥有的臻享卡数量 API请求
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
type AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripmerchantgalaxyderbymembervoucherqueryamountRequest 初始化AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherqueryamountRequest() *AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.query.amount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest) GetToken() string {
	return r._token
}
