package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取小程序分享文案和图片 APIRequest
alitrip.merchant.galaxy.share.get

获取 雅高微信小程序分享素材文案和图片。
*/
type AlitripMerchantGalaxyShareGetRequest struct {
    model.Params

    // 租户ID
    tenantKey   string 

}

func NewAlitripMerchantGalaxyShareGetRequest() *AlitripMerchantGalaxyShareGetRequest{
    return &AlitripMerchantGalaxyShareGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyShareGetRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.share.get"
}

func (r AlitripMerchantGalaxyShareGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyShareGetRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyShareGetRequest) GetTenantKey() string {
    return r.tenantKey
}

