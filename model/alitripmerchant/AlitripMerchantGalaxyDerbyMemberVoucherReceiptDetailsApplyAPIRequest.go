package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest v5.0付费会员卡开发订单开票详情申请 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply
//
// v5.0付费会员卡开发订单开票详情申请
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 开票参数
	_derbyVoucherCardApplyReceiptDTO *DerbyVoucherCardApplyReceiptDto
}

// NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyRequest 初始化AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyRequest() *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardApplyReceiptDTO is DerbyVoucherCardApplyReceiptDTO Setter
// 开票参数
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) SetDerbyVoucherCardApplyReceiptDTO(_derbyVoucherCardApplyReceiptDTO *DerbyVoucherCardApplyReceiptDto) error {
	r._derbyVoucherCardApplyReceiptDTO = _derbyVoucherCardApplyReceiptDTO
	r.Set("derby_voucher_card_apply_receipt_d_t_o", _derbyVoucherCardApplyReceiptDTO)
	return nil
}

// GetDerbyVoucherCardApplyReceiptDTO DerbyVoucherCardApplyReceiptDTO Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest) GetDerbyVoucherCardApplyReceiptDTO() *DerbyVoucherCardApplyReceiptDto {
	return r._derbyVoucherCardApplyReceiptDTO
}
