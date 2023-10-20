package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest v5.0付费会员卡开发订单开票详情申请 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply
//
// v5.0付费会员卡开发订单开票详情申请
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 开票参数
	_derbyVoucherCardApplyReceiptDTO *DerbyVoucherCardApplyReceiptDto
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._derbyVoucherCardApplyReceiptDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardApplyReceiptDTO is DerbyVoucherCardApplyReceiptDTO Setter
// 开票参数
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) SetDerbyVoucherCardApplyReceiptDTO(_derbyVoucherCardApplyReceiptDTO *DerbyVoucherCardApplyReceiptDto) error {
	r._derbyVoucherCardApplyReceiptDTO = _derbyVoucherCardApplyReceiptDTO
	r.Set("derby_voucher_card_apply_receipt_d_t_o", _derbyVoucherCardApplyReceiptDTO)
	return nil
}

// GetDerbyVoucherCardApplyReceiptDTO DerbyVoucherCardApplyReceiptDTO Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) GetDerbyVoucherCardApplyReceiptDTO() *DerbyVoucherCardApplyReceiptDto {
	return r._derbyVoucherCardApplyReceiptDTO
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyAPIRequest.Put(v)
}
