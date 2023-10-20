package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest 德比付费会员卡开票自匹配 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.auto
//
// 德比付费会员卡开票自匹配
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 关键词
	_keyword string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) Reset() {
	r._tenantKey = ""
	r._keyword = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.auto"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) GetKeyword() string {
	return r._keyword
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest.Put(v)
}
