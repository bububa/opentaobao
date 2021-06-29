package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-用户登出 API请求
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

// 初始化AlitripMerchantGalaxyMemberLogoutRequest对象
func NewAlitripMerchantGalaxyMemberLogoutRequest() *AlitripMerchantGalaxyMemberLogoutRequest{
    return &AlitripMerchantGalaxyMemberLogoutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberLogoutRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.logout"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberLogoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyMemberLogoutRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberLogoutRequest) GetTenantKey() string {
    return r.tenantKey
}
// Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyMemberLogoutRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyMemberLogoutRequest) GetToken() string {
    return r.token
}
