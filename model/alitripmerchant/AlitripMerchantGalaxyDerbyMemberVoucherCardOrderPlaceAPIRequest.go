package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest 德比付费会员卡下单 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
type AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 商品参数
	_derbyVoucherCardPayDTO *DerbyVoucherCardPayDto
}

// NewAlitripmerchantgalaxyderbymembervouchercardorderplaceRequest 初始化AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardorderplaceRequest() *AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.place"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardPayDTO is DerbyVoucherCardPayDTO Setter
// 商品参数
func (r *AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) SetDerbyVoucherCardPayDTO(_derbyVoucherCardPayDTO *DerbyVoucherCardPayDto) error {
	r._derbyVoucherCardPayDTO = _derbyVoucherCardPayDTO
	r.Set("derby_voucher_card_pay_d_t_o", _derbyVoucherCardPayDTO)
	return nil
}

// GetDerbyVoucherCardPayDTO DerbyVoucherCardPayDTO Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest) GetDerbyVoucherCardPayDTO() *DerbyVoucherCardPayDto {
	return r._derbyVoucherCardPayDTO
}
