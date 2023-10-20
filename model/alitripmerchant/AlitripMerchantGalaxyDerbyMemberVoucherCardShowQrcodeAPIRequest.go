package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest 会员权益卡身份识别二维码图片 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest struct {
	model.Params
	// 123
	_tenantKey string
	// 100
	_token string
	// 卡号
	_cardNumber string
	// 权益卡号
	_memberVoucherCardID string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._cardNumber = ""
	r._memberVoucherCardID = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 123
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 100
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetCardNumber is CardNumber Setter
// 卡号
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) SetCardNumber(_cardNumber string) error {
	r._cardNumber = _cardNumber
	r.Set("card_number", _cardNumber)
	return nil
}

// GetCardNumber CardNumber Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetCardNumber() string {
	return r._cardNumber
}

// SetMemberVoucherCardID is MemberVoucherCardID Setter
// 权益卡号
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) SetMemberVoucherCardID(_memberVoucherCardID string) error {
	r._memberVoucherCardID = _memberVoucherCardID
	r.Set("member_voucher_card_i_d", _memberVoucherCardID)
	return nil
}

// GetMemberVoucherCardID MemberVoucherCardID Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) GetMemberVoucherCardID() string {
	return r._memberVoucherCardID
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest.Put(v)
}
