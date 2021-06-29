package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取会员卡信息 APIRequest
alitrip.merchant.galaxy.member.card

星河=根据会员等级获取会员的权益
*/
type AlitripMerchantGalaxyMemberCardRequest struct {
    model.Params

    // 租户信息
    tenantKey   string 

    // 飞猪等级
    fliggyLevel   string 

    // 卡类型
    cardType   string 

}

func NewAlitripMerchantGalaxyMemberCardRequest() *AlitripMerchantGalaxyMemberCardRequest{
    return &AlitripMerchantGalaxyMemberCardRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberCardRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.card"
}

func (r AlitripMerchantGalaxyMemberCardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberCardRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberCardRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberCardRequest) SetFliggyLevel(fliggyLevel string) error {
    r.fliggyLevel = fliggyLevel
    r.Set("fliggy_level", fliggyLevel)
    return nil
}

func (r AlitripMerchantGalaxyMemberCardRequest) GetFliggyLevel() string {
    return r.fliggyLevel
}

func (r *AlitripMerchantGalaxyMemberCardRequest) SetCardType(cardType string) error {
    r.cardType = cardType
    r.Set("card_type", cardType)
    return nil
}

func (r AlitripMerchantGalaxyMemberCardRequest) GetCardType() string {
    return r.cardType
}

