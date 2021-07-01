package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外提供会员注册服务 API请求
alitrip.merchant.galaxy.member.provider.register

星河产品=对外提供注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberProviderRegisterAPIRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // 注册入参
    _providerMemberParam   *ProviderMemberParam
}

// 初始化AlitripMerchantGalaxyMemberProviderRegisterAPIRequest对象
func NewAlitripMerchantGalaxyMemberProviderRegisterRequest() *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest{
    return &AlitripMerchantGalaxyMemberProviderRegisterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.provider.register"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// ProviderMemberParam Setter
// 注册入参
func (r *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) SetProviderMemberParam(_providerMemberParam *ProviderMemberParam) error {
    r._providerMemberParam = _providerMemberParam
    r.Set("provider_member_param", _providerMemberParam)
    return nil
}

// ProviderMemberParam Getter
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetProviderMemberParam() *ProviderMemberParam {
    return r._providerMemberParam
}
