package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取会员卡信息 API请求
alitrip.merchant.galaxy.member.card

星河=根据会员等级获取会员的权益
*/
type AlitripMerchantGalaxyMemberCardRequest struct {
    model.Params
    // 租户信息
    _tenantKey   string
    // 飞猪等级
    _fliggyLevel   string
    // 卡类型
    _cardType   string
}

// 初始化AlitripMerchantGalaxyMemberCardRequest对象
func NewAlitripMerchantGalaxyMemberCardRequest() *AlitripMerchantGalaxyMemberCardRequest{
    return &AlitripMerchantGalaxyMemberCardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberCardRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.card"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberCardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyMemberCardRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberCardRequest) GetTenantKey() string {
    return r._tenantKey
}
// FliggyLevel Setter
// 飞猪等级
func (r *AlitripMerchantGalaxyMemberCardRequest) SetFliggyLevel(_fliggyLevel string) error {
    r._fliggyLevel = _fliggyLevel
    r.Set("fliggy_level", _fliggyLevel)
    return nil
}

// FliggyLevel Getter
func (r AlitripMerchantGalaxyMemberCardRequest) GetFliggyLevel() string {
    return r._fliggyLevel
}
// CardType Setter
// 卡类型
func (r *AlitripMerchantGalaxyMemberCardRequest) SetCardType(_cardType string) error {
    r._cardType = _cardType
    r.Set("card_type", _cardType)
    return nil
}

// CardType Getter
func (r AlitripMerchantGalaxyMemberCardRequest) GetCardType() string {
    return r._cardType
}
