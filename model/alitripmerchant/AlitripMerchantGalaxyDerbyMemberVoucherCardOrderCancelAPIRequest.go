package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest 取消订单 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 取消订单参数
	_derbyVoucherCardCancelDTO *DerbyVoucherCardCancelDto
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardCancelDTO is DerbyVoucherCardCancelDTO Setter
// 取消订单参数
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) SetDerbyVoucherCardCancelDTO(_derbyVoucherCardCancelDTO *DerbyVoucherCardCancelDto) error {
	r._derbyVoucherCardCancelDTO = _derbyVoucherCardCancelDTO
	r.Set("derby_voucher_card_cancel_d_t_o", _derbyVoucherCardCancelDTO)
	return nil
}

// GetDerbyVoucherCardCancelDTO DerbyVoucherCardCancelDTO Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest) GetDerbyVoucherCardCancelDTO() *DerbyVoucherCardCancelDto {
	return r._derbyVoucherCardCancelDTO
}
