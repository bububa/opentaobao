package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-微信小程序会员注册 APIRequest
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberRegisterRequest struct {
    model.Params

    // 租户身份信息
    tenantKey   string 

    // 注册入参
    memberParam   *MemberParam 

}

func NewAlitripMerchantGalaxyMemberRegisterRequest() *AlitripMerchantGalaxyMemberRegisterRequest{
    return &AlitripMerchantGalaxyMemberRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyMemberRegisterRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.register"
}

func (r AlitripMerchantGalaxyMemberRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyMemberRegisterRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyMemberRegisterRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyMemberRegisterRequest) SetMemberParam(memberParam *MemberParam) error {
    r.memberParam = memberParam
    r.Set("member_param", memberParam)
    return nil
}

func (r AlitripMerchantGalaxyMemberRegisterRequest) GetMemberParam() *MemberParam {
    return r.memberParam
}

