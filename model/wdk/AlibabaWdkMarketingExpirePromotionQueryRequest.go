package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠查询 APIRequest
alibaba.wdk.marketing.expire.promotion.query

短保优惠查询
*/
type AlibabaWdkMarketingExpirePromotionQueryRequest struct {
    model.Params

    // 店铺id
    shopId   string 

    // 商品skucode
    skuCode   string 

}

func NewAlibabaWdkMarketingExpirePromotionQueryRequest() *AlibabaWdkMarketingExpirePromotionQueryRequest{
    return &AlibabaWdkMarketingExpirePromotionQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.query"
}

func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetShopId() string {
    return r.shopId
}

func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetSkuCode() string {
    return r.skuCode
}

