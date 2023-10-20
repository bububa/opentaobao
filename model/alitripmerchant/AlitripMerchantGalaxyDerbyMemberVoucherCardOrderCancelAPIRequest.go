package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest 取消订单 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
type AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 取消订单参数
	_derbyVoucherCardCancelDTO *DerbyVoucherCardCancelDto
}

// NewAlitripmerchantgalaxyderbymembervouchercardordercancelRequest 初始化AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardordercancelRequest() *AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardCancelDTO is DerbyVoucherCardCancelDTO Setter
// 取消订单参数
func (r *AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) SetDerbyVoucherCardCancelDTO(_derbyVoucherCardCancelDTO *DerbyVoucherCardCancelDto) error {
	r._derbyVoucherCardCancelDTO = _derbyVoucherCardCancelDTO
	r.Set("derby_voucher_card_cancel_d_t_o", _derbyVoucherCardCancelDTO)
	return nil
}

// GetDerbyVoucherCardCancelDTO DerbyVoucherCardCancelDTO Getter
func (r AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest) GetDerbyVoucherCardCancelDTO() *DerbyVoucherCardCancelDto {
	return r._derbyVoucherCardCancelDTO
}
