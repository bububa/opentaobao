package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取微信用户的信息 API请求
alitrip.merchant.galaxy.wechat.info

获取微信用户的openId和unionId
*/
type AlitripMerchantGalaxyWechatInfoRequest struct {
    model.Params
    // 租户的id
    tenantKey   string
    // 微信小程序获取的code
    code   string
}

// 初始化AlitripMerchantGalaxyWechatInfoRequest对象
func NewAlitripMerchantGalaxyWechatInfoRequest() *AlitripMerchantGalaxyWechatInfoRequest{
    return &AlitripMerchantGalaxyWechatInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatInfoRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.wechat.info"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户的id
func (r *AlitripMerchantGalaxyWechatInfoRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyWechatInfoRequest) GetTenantKey() string {
    return r.tenantKey
}
// Code Setter
// 微信小程序获取的code
func (r *AlitripMerchantGalaxyWechatInfoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlitripMerchantGalaxyWechatInfoRequest) GetCode() string {
    return r.code
}
