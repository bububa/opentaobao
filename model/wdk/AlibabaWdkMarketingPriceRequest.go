package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
促销价签服务 APIRequest
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

func NewAlibabaWdkMarketingPriceRequest() *AlibabaWdkMarketingPriceRequest{
    return &AlibabaWdkMarketingPriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingPriceRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.price"
}

func (r AlibabaWdkMarketingPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingPriceRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaWdkMarketingPriceRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *AlibabaWdkMarketingPriceRequest) SetSkuCodes(skuCodes []string) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetSkuCodes() []string {
    return r.skuCodes
}

func (r *AlibabaWdkMarketingPriceRequest) SetShopIds(shopIds []int64) error {
    r.shopIds = shopIds
    r.Set("shop_ids", shopIds)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetShopIds() []int64 {
    return r.shopIds
}

func (r *AlibabaWdkMarketingPriceRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetEndTime() string {
    return r.endTime
}

func (r *AlibabaWdkMarketingPriceRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r AlibabaWdkMarketingPriceRequest) GetBeginTime() string {
    return r.beginTime
}

