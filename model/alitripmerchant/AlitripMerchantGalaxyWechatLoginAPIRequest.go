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
type AlitripMerchantGalaxyWechatLoginAPIRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // 微信小程序登陆请求参数
    _loginParam   *LoginParam
}

// 初始化AlitripMerchantGalaxyWechatLoginAPIRequest对象
func NewAlitripMerchantGalaxyWechatLoginRequest() *AlitripMerchantGalaxyWechatLoginAPIRequest{
    return &AlitripMerchantGalaxyWechatLoginAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.wechat.login"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyWechatLoginAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// LoginParam Setter
// 微信小程序登陆请求参数
func (r *AlitripMerchantGalaxyWechatLoginAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
    r._loginParam = _loginParam
    r.Set("login_param", _loginParam)
    return nil
}

// LoginParam Getter
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetLoginParam() *LoginParam {
    return r._loginParam
}
