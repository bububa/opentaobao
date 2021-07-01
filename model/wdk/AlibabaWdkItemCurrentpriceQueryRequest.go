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
type AlibabaWdkItemCurrentpriceQueryAPIRequest struct {
    model.Params
    // 渠道店id
    _shopId   int64
    // sku编码列表
    _skuCodes   []string
    // 渠道
    _orderChannelCode   string
}

// 初始化AlibabaWdkItemCurrentpriceQueryAPIRequest对象
func NewAlibabaWdkItemCurrentpriceQueryRequest() *AlibabaWdkItemCurrentpriceQueryAPIRequest{
    return &AlibabaWdkItemCurrentpriceQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.currentprice.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetShopId() int64 {
    return r._shopId
}
// SkuCodes Setter
// sku编码列表
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetSkuCodes(_skuCodes []string) error {
    r._skuCodes = _skuCodes
    r.Set("sku_codes", _skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetSkuCodes() []string {
    return r._skuCodes
}
// OrderChannelCode Setter
// 渠道
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetOrderChannelCode(_orderChannelCode string) error {
    r._orderChannelCode = _orderChannelCode
    r.Set("order_channel_code", _orderChannelCode)
    return nil
}

// OrderChannelCode Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetOrderChannelCode() string {
    return r._orderChannelCode
}
