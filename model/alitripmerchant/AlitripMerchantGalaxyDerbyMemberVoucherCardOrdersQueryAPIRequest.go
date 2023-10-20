package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest 查询权益卡订单列表 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
type AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 状态
	_status string
}

// NewAlitripmerchantgalaxyderbymembervouchercardordersqueryRequest 初始化AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardordersqueryRequest() *AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.orders.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetToken() string {
	return r._token
}

// SetStatus is Status Setter
// 状态
func (r *AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest) GetStatus() string {
	return r._status
}
