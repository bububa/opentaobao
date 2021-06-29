package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个商品未来价查询接口 APIRequest
alibaba.wdk.item.futureprice.query

查询单个商品未来价，融合了未来基础售价+未来促销价
*/
type AlibabaWdkItemFuturepriceQueryRequest struct {
    model.Params

    // 渠道店id
    shopId   int64 

    // 商品编码
    skuCode   string 

    // 渠道
    orderChannelCode   string 

    // 开始时间
    startTime   string 

    // 结束时间，结束时间-开始时间不能超过48小时
    endTime   string 

}

func NewAlibabaWdkItemFuturepriceQueryRequest() *AlibabaWdkItemFuturepriceQueryRequest{
    return &AlibabaWdkItemFuturepriceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.futureprice.query"
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemFuturepriceQueryRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetShopId() int64 {
    return r.shopId
}

func (r *AlibabaWdkItemFuturepriceQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkItemFuturepriceQueryRequest) SetOrderChannelCode(orderChannelCode string) error {
    r.orderChannelCode = orderChannelCode
    r.Set("order_channel_code", orderChannelCode)
    return nil
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetOrderChannelCode() string {
    return r.orderChannelCode
}

func (r *AlibabaWdkItemFuturepriceQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetStartTime() string {
    return r.startTime
}

func (r *AlibabaWdkItemFuturepriceQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaWdkItemFuturepriceQueryRequest) GetEndTime() string {
    return r.endTime
}

