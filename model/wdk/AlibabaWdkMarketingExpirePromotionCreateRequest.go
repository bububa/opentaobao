package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠创建 APIRequest
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
type AlibabaWdkMarketingExpirePromotionCreateRequest struct {
    model.Params

    // 创建短保优惠
    param0   *ExpirePromotionBo 

}

func NewAlibabaWdkMarketingExpirePromotionCreateRequest() *AlibabaWdkMarketingExpirePromotionCreateRequest{
    return &AlibabaWdkMarketingExpirePromotionCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.create"
}

func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingExpirePromotionCreateRequest) SetParam0(param0 *ExpirePromotionBo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetParam0() *ExpirePromotionBo {
    return r.param0
}

