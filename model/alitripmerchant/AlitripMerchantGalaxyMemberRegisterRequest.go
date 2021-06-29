package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-微信小程序会员注册 API请求
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberRegisterRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // 注册入参
    _memberParam   *MemberParam
}

// 初始化AlitripMerchantGalaxyMemberRegisterRequest对象
func NewAlitripMerchantGalaxyMemberRegisterRequest() *AlitripMerchantGalaxyMemberRegisterRequest{
    return &AlitripMerchantGalaxyMemberRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberRegisterRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.register"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberRegisterRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberRegisterRequest) GetTenantKey() string {
    return r._tenantKey
}
// MemberParam Setter
// 注册入参
func (r *AlitripMerchantGalaxyMemberRegisterRequest) SetMemberParam(_memberParam *MemberParam) error {
    r._memberParam = _memberParam
    r.Set("member_param", _memberParam)
    return nil
}

// MemberParam Getter
func (r AlitripMerchantGalaxyMemberRegisterRequest) GetMemberParam() *MemberParam {
    return r._memberParam
}
