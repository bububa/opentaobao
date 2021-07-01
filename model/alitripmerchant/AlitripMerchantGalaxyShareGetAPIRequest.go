package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取小程序分享文案和图片 API请求
alitrip.merchant.galaxy.share.get

获取 雅高微信小程序分享素材文案和图片。
*/
type AlitripMerchantGalaxyShareGetAPIRequest struct {
    model.Params
    // 租户ID
    _tenantKey   string
}

// 初始化AlitripMerchantGalaxyShareGetAPIRequest对象
func NewAlitripMerchantGalaxyShareGetRequest() *AlitripMerchantGalaxyShareGetAPIRequest{
    return &AlitripMerchantGalaxyShareGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.share.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyShareGetAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyShareGetAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
