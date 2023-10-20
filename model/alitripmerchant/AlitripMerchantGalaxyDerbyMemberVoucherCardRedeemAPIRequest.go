package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest 根据兑换码兑换臻享卡接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
type AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户标识
	_token string
	// 兑换臻享卡入参
	_derbyVoucherCardRedeemDTO *DerbyVoucherCardRedeemDto
}

// NewAlitripmerchantgalaxyderbymembervouchercardredeemRequest 初始化AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardredeemRequest() *AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.redeem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardRedeemDTO is DerbyVoucherCardRedeemDTO Setter
// 兑换臻享卡入参
func (r *AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) SetDerbyVoucherCardRedeemDTO(_derbyVoucherCardRedeemDTO *DerbyVoucherCardRedeemDto) error {
	r._derbyVoucherCardRedeemDTO = _derbyVoucherCardRedeemDTO
	r.Set("derby_voucher_card_redeem_d_t_o", _derbyVoucherCardRedeemDTO)
	return nil
}

// GetDerbyVoucherCardRedeemDTO DerbyVoucherCardRedeemDTO Getter
func (r AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest) GetDerbyVoucherCardRedeemDTO() *DerbyVoucherCardRedeemDto {
	return r._derbyVoucherCardRedeemDTO
}
