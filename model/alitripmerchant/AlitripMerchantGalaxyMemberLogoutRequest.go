package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-用户登出 APIRequest
alitrip.merchant.galaxy.member.logout

星河=微信小程序用户登出
*/
type AlitripMerchantGalaxyMemberLogoutRequest struct {
    model.Params

    // 租户信息
    tenantKey   string 

    // 用户登录token
    token   string 

}

func NewAlitripMerchantGalaxyMemberLogoutRequest() *AlitripMerchantGalaxyMemberLogoutRequest{
    return &AlitripMerchantGalaxyMemberLogoutRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberLogoutRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.logout"
}

func (r AlitripMerchantGalaxyMemberLogoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberLogoutRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberLogoutRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberLogoutRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyMemberLogoutRequest) GetToken() string {
    return r.token
}

