package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品当前价格 API请求
alibaba.wdk.item.currentprice.query

通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量<=20,返回结果key为skuCode
*/
type AlibabaWdkItemCurrentpriceQueryRequest struct {
    model.Params
    // 渠道店id
    shopId   int64
    // sku编码列表
    skuCodes   []string
    // 渠道
    orderChannelCode   string
}

// 初始化AlibabaWdkItemCurrentpriceQueryRequest对象
func NewAlibabaWdkItemCurrentpriceQueryRequest() *AlibabaWdkItemCurrentpriceQueryRequest{
    return &AlibabaWdkItemCurrentpriceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCurrentpriceQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.currentprice.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCurrentpriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemCurrentpriceQueryRequest) GetShopId() int64 {
    return r.shopId
}
// SkuCodes Setter
// sku编码列表
func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetSkuCodes(skuCodes []string) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaWdkItemCurrentpriceQueryRequest) GetSkuCodes() []string {
    return r.skuCodes
}
// OrderChannelCode Setter
// 渠道
func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetOrderChannelCode(orderChannelCode string) error {
    r.orderChannelCode = orderChannelCode
    r.Set("order_channel_code", orderChannelCode)
    return nil
}

// OrderChannelCode Getter
func (r AlibabaWdkItemCurrentpriceQueryRequest) GetOrderChannelCode() string {
    return r.orderChannelCode
}
