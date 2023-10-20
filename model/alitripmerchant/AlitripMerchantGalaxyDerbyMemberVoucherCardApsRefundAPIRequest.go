package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest Aps退券通知接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
type AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 请求String
	_derbyBody string
}

// NewAlitripmerchantgalaxyderbymembervouchercardapsrefundRequest 初始化AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardapsrefundRequest() *AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDerbyBody is DerbyBody Setter
// 请求String
func (r *AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) SetDerbyBody(_derbyBody string) error {
	r._derbyBody = _derbyBody
	r.Set("derby_body", _derbyBody)
	return nil
}

// GetDerbyBody DerbyBody Getter
func (r AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest) GetDerbyBody() string {
	return r._derbyBody
}
