package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠删除 APIRequest
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除
*/
type AlibabaWdkMarketingExpirePromotionDeleteRequest struct {
    model.Params

    // 删除短保优惠
    param0   *ExpirePromotionBo 

}

func NewAlibabaWdkMarketingExpirePromotionDeleteRequest() *AlibabaWdkMarketingExpirePromotionDeleteRequest{
    return &AlibabaWdkMarketingExpirePromotionDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.delete"
}

func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingExpirePromotionDeleteRequest) SetParam0(param0 *ExpirePromotionBo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetParam0() *ExpirePromotionBo {
    return r.param0
}

