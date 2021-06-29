package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-用户使用微信登陆 API请求
alitrip.merchant.galaxy.wechat.login

星河产品=用户微信小程序登陆
*/
type AlitripMerchantGalaxyWechatLoginRequest struct {
    model.Params
    // 租户身份信息
    tenantKey   string
    // 微信小程序登陆请求参数
    loginParam   *LoginParam
}

// 初始化AlitripMerchantGalaxyWechatLoginRequest对象
func NewAlitripMerchantGalaxyWechatLoginRequest() *AlitripMerchantGalaxyWechatLoginRequest{
    return &AlitripMerchantGalaxyWechatLoginRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatLoginRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.wechat.login"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyWechatLoginRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyWechatLoginRequest) GetTenantKey() string {
    return r.tenantKey
}
// LoginParam Setter
// 微信小程序登陆请求参数
func (r *AlitripMerchantGalaxyWechatLoginRequest) SetLoginParam(loginParam *LoginParam) error {
    r.loginParam = loginParam
    r.Set("login_param", loginParam)
    return nil
}

// LoginParam Getter
func (r AlitripMerchantGalaxyWechatLoginRequest) GetLoginParam() *LoginParam {
    return r.loginParam
}
