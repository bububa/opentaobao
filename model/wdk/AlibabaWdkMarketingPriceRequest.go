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
    _pageSize   int64
    // 页码
    _pageIndex   int64
    // 商品sku
    _skuCodes   []string
    // 门店标识数组
    _shopIds   []int64
    // 查询结束时间(sku_codes非空无效)
    _endTime   string
    // 查询开始时间(sku_codes非空无效)
    _beginTime   string
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
func (r *AlibabaWdkMarketingPriceRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkMarketingPriceRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageIndex Setter
// 页码
func (r *AlibabaWdkMarketingPriceRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AlibabaWdkMarketingPriceRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// SkuCodes Setter
// 商品sku
func (r *AlibabaWdkMarketingPriceRequest) SetSkuCodes(_skuCodes []string) error {
    r._skuCodes = _skuCodes
    r.Set("sku_codes", _skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaWdkMarketingPriceRequest) GetSkuCodes() []string {
    return r._skuCodes
}
// ShopIds Setter
// 门店标识数组
func (r *AlibabaWdkMarketingPriceRequest) SetShopIds(_shopIds []int64) error {
    r._shopIds = _shopIds
    r.Set("shop_ids", _shopIds)
    return nil
}

// ShopIds Getter
func (r AlibabaWdkMarketingPriceRequest) GetShopIds() []int64 {
    return r._shopIds
}
// EndTime Setter
// 查询结束时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWdkMarketingPriceRequest) GetEndTime() string {
    return r._endTime
}
// BeginTime Setter
// 查询开始时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceRequest) SetBeginTime(_beginTime string) error {
    r._beginTime = _beginTime
    r.Set("begin_time", _beginTime)
    return nil
}

// BeginTime Getter
func (r AlibabaWdkMarketingPriceRequest) GetBeginTime() string {
    return r._beginTime
}
