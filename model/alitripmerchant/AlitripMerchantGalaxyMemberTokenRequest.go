package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-校验token APIRequest
alitrip.merchant.galaxy.member.token

校验或者刷新token
*/
type AlitripMerchantGalaxyMemberTokenRequest struct {
    model.Params

    // 租户身份信息
    tenantKey   string 

    // 访问携带toekn
    token   string 

}

func NewAlitripMerchantGalaxyMemberTokenRequest() *AlitripMerchantGalaxyMemberTokenRequest{
    return &AlitripMerchantGalaxyMemberTokenRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberTokenRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.token"
}

func (r AlitripMerchantGalaxyMemberTokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberTokenRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberTokenRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberTokenRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyMemberTokenRequest) GetToken() string {
    return r.token
}

