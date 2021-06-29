package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠查询 API请求
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

// 初始化AlibabaWdkMarketingExpirePromotionQueryRequest对象
func NewAlibabaWdkMarketingExpirePromotionQueryRequest() *AlibabaWdkMarketingExpirePromotionQueryRequest{
    return &AlibabaWdkMarketingExpirePromotionQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 店铺id
func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetShopId() string {
    return r.shopId
}
// SkuCode Setter
// 商品skucode
func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetSkuCode() string {
    return r.skuCode
}
