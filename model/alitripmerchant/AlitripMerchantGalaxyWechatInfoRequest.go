package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取微信用户的信息 APIRequest
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

func NewAlitripMerchantGalaxyWechatInfoRequest() *AlitripMerchantGalaxyWechatInfoRequest{
    return &AlitripMerchantGalaxyWechatInfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyWechatInfoRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.wechat.info"
}

func (r AlitripMerchantGalaxyWechatInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyWechatInfoRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyWechatInfoRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyWechatInfoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlitripMerchantGalaxyWechatInfoRequest) GetCode() string {
    return r.code
}

