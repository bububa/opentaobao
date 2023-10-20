package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest 德比付费会员卡下单 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 商品参数
	_derbyVoucherCardPayDTO *DerbyVoucherCardPayDto
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._derbyVoucherCardPayDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.place"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetToken() string {
	return r._token
}

// SetDerbyVoucherCardPayDTO is DerbyVoucherCardPayDTO Setter
// 商品参数
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) SetDerbyVoucherCardPayDTO(_derbyVoucherCardPayDTO *DerbyVoucherCardPayDto) error {
	r._derbyVoucherCardPayDTO = _derbyVoucherCardPayDTO
	r.Set("derby_voucher_card_pay_d_t_o", _derbyVoucherCardPayDTO)
	return nil
}

// GetDerbyVoucherCardPayDTO DerbyVoucherCardPayDTO Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) GetDerbyVoucherCardPayDTO() *DerbyVoucherCardPayDto {
	return r._derbyVoucherCardPayDTO
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceAPIRequest.Put(v)
}
