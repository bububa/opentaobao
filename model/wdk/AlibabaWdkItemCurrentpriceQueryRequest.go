package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品当前价格 APIRequest
alibaba.wdk.item.currentprice.query

通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量<=20,返回结果key为skuCode
*/
type AlibabaWdkItemCurrentpriceQueryRequest struct {
    model.Params

    // 渠道店id
    shopId   int64 

    // sku编码列表
    skuCodes   []String 

    // 渠道
    orderChannelCode   string 

}

func NewAlibabaWdkItemCurrentpriceQueryRequest() *AlibabaWdkItemCurrentpriceQueryRequest{
    return &AlibabaWdkItemCurrentpriceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemCurrentpriceQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.currentprice.query"
}

func (r AlibabaWdkItemCurrentpriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkItemCurrentpriceQueryRequest) GetShopId() int64 {
    return r.shopId
}

func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetSkuCodes(skuCodes []String) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

func (r AlibabaWdkItemCurrentpriceQueryRequest) GetSkuCodes() []String {
    return r.skuCodes
}

func (r *AlibabaWdkItemCurrentpriceQueryRequest) SetOrderChannelCode(orderChannelCode string) error {
    r.orderChannelCode = orderChannelCode
    r.Set("order_channel_code", orderChannelCode)
    return nil
}

func (r AlibabaWdkItemCurrentpriceQueryRequest) GetOrderChannelCode() string {
    return r.orderChannelCode
}

