package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外提供会员注册服务 APIRequest
alitrip.merchant.galaxy.member.provider.register

星河产品=对外提供注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberProviderRegisterRequest struct {
    model.Params

    // 租户身份信息
    tenantKey   string 

    // 注册入参
    providerMemberParam   *ProviderMemberParam 

}

func NewAlitripMerchantGalaxyMemberProviderRegisterRequest() *AlitripMerchantGalaxyMemberProviderRegisterRequest{
    return &AlitripMerchantGalaxyMemberProviderRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberProviderRegisterRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.provider.register"
}

func (r AlitripMerchantGalaxyMemberProviderRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberProviderRegisterRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberProviderRegisterRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberProviderRegisterRequest) SetProviderMemberParam(providerMemberParam *ProviderMemberParam) error {
    r.providerMemberParam = providerMemberParam
    r.Set("provider_member_param", providerMemberParam)
    return nil
}

func (r AlitripMerchantGalaxyMemberProviderRegisterRequest) GetProviderMemberParam() *ProviderMemberParam {
    return r.providerMemberParam
}

