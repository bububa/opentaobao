package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取登录用户的信息 APIRequest
alitrip.merchant.galaxy.member.query

获取登录用户的信息
*/
type AlitripMerchantGalaxyMemberQueryRequest struct {
    model.Params

    // 租户身份信息
    tenantKey   string 

    // toekn
    token   string 

}

func NewAlitripMerchantGalaxyMemberQueryRequest() *AlitripMerchantGalaxyMemberQueryRequest{
    return &AlitripMerchantGalaxyMemberQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.query"
}

func (r AlitripMerchantGalaxyMemberQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberQueryRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyMemberQueryRequest) GetToken() string {
    return r.token
}

