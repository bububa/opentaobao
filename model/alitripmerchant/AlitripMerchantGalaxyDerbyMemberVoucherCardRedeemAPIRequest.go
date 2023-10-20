package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest 根据兑换码兑换臻享卡接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户标识
	_token string
	// 兑换臻享卡入参
	_derbyVoucherCardRedeemDTO *DerbyVoucherCardRedeemDto
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._derbyVoucherCardRedeemDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.redeem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardRedeemDTO is DerbyVoucherCardRedeemDTO Setter
// 兑换臻享卡入参
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) SetDerbyVoucherCardRedeemDTO(_derbyVoucherCardRedeemDTO *DerbyVoucherCardRedeemDto) error {
	r._derbyVoucherCardRedeemDTO = _derbyVoucherCardRedeemDTO
	r.Set("derby_voucher_card_redeem_d_t_o", _derbyVoucherCardRedeemDTO)
	return nil
}

// GetDerbyVoucherCardRedeemDTO DerbyVoucherCardRedeemDTO Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) GetDerbyVoucherCardRedeemDTO() *DerbyVoucherCardRedeemDto {
	return r._derbyVoucherCardRedeemDTO
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest.Put(v)
}
