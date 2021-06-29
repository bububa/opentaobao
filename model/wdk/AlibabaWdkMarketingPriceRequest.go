package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
促销价签服务 API请求
alibaba.wdk.marketing.price

获取营销-促销商品中的实时价格
*/
type AlibabaWdkMarketingPriceRequest struct {
    model.Params
    // 单页大小
    pageSize   int64
    // 页码
    pageIndex   int64
    // 商品sku
    skuCodes   []string
    // 门店标识数组
    shopIds   []int64
    // 查询结束时间(sku_codes非空无效)
    endTime   string
    // 查询开始时间(sku_codes非空无效)
    beginTime   string
}

// 初始化AlibabaWdkMarketingPriceRequest对象
func NewAlibabaWdkMarketingPriceRequest() *AlibabaWdkMarketingPriceRequest{
    return &AlibabaWdkMarketingPriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingPriceRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.price"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 单页大小
func (r *AlibabaWdkMarketingPriceRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkMarketingPriceRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageIndex Setter
// 页码
func (r *AlibabaWdkMarketingPriceRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r AlibabaWdkMarketingPriceRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// SkuCodes Setter
// 商品sku
func (r *AlibabaWdkMarketingPriceRequest) SetSkuCodes(skuCodes []string) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaWdkMarketingPriceRequest) GetSkuCodes() []string {
    return r.skuCodes
}
// ShopIds Setter
// 门店标识数组
func (r *AlibabaWdkMarketingPriceRequest) SetShopIds(shopIds []int64) error {
    r.shopIds = shopIds
    r.Set("shop_ids", shopIds)
    return nil
}

// ShopIds Getter
func (r AlibabaWdkMarketingPriceRequest) GetShopIds() []int64 {
    return r.shopIds
}
// EndTime Setter
// 查询结束时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWdkMarketingPriceRequest) GetEndTime() string {
    return r.endTime
}
// BeginTime Setter
// 查询开始时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

// BeginTime Getter
func (r AlibabaWdkMarketingPriceRequest) GetBeginTime() string {
    return r.beginTime
}
