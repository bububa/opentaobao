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
    _shopId   string
    // 商品skucode
    _skuCode   string
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
func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetShopId() string {
    return r._shopId
}
// SkuCode Setter
// 商品skucode
func (r *AlibabaWdkMarketingExpirePromotionQueryRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMarketingExpirePromotionQueryRequest) GetSkuCode() string {
    return r._skuCode
}
