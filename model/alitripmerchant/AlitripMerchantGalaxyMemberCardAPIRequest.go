package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberCardAPIRequest 星河-获取会员卡信息 API请求
// alitrip.merchant.galaxy.member.card
//
// 星河=根据会员等级获取会员的权益
type AlitripMerchantGalaxyMemberCardAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
	// 飞猪等级
	_fliggyLevel string
	// 卡类型
	_cardType string
}

// NewAlitripMerchantGalaxyMemberCardRequest 初始化AlitripMerchantGalaxyMemberCardAPIRequest对象
func NewAlitripMerchantGalaxyMemberCardRequest() *AlitripMerchantGalaxyMemberCardAPIRequest {
	return &AlitripMerchantGalaxyMemberCardAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberCardAPIRequest) Reset() {
	r._tenantKey = ""
	r._fliggyLevel = ""
	r._cardType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.card"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyMemberCardAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetFliggyLevel is FliggyLevel Setter
// 飞猪等级
func (r *AlitripMerchantGalaxyMemberCardAPIRequest) SetFliggyLevel(_fliggyLevel string) error {
	r._fliggyLevel = _fliggyLevel
	r.Set("fliggy_level", _fliggyLevel)
	return nil
}

// GetFliggyLevel FliggyLevel Getter
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetFliggyLevel() string {
	return r._fliggyLevel
}

// SetCardType is CardType Setter
// 卡类型
func (r *AlitripMerchantGalaxyMemberCardAPIRequest) SetCardType(_cardType string) error {
	r._cardType = _cardType
	r.Set("card_type", _cardType)
	return nil
}

// GetCardType CardType Getter
func (r AlitripMerchantGalaxyMemberCardAPIRequest) GetCardType() string {
	return r._cardType
}

var poolAlitripMerchantGalaxyMemberCardAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberCardRequest()
	},
}

// GetAlitripMerchantGalaxyMemberCardRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberCardAPIRequest
func GetAlitripMerchantGalaxyMemberCardAPIRequest() *AlitripMerchantGalaxyMemberCardAPIRequest {
	return poolAlitripMerchantGalaxyMemberCardAPIRequest.Get().(*AlitripMerchantGalaxyMemberCardAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberCardAPIRequest 将 AlitripMerchantGalaxyMemberCardAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberCardAPIRequest(v *AlitripMerchantGalaxyMemberCardAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberCardAPIRequest.Put(v)
}
