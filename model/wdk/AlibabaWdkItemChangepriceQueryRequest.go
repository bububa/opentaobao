package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照价格变更时间段，查询会变更价格的单据的商品 API请求
alibaba.wdk.item.changeprice.query

/**
     * 按照价格变更时间段，查询会变更价格的单据的商品
     * 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
     * 返回活动开始时间在 startTime<=活动开始时间<endTime 之间的所有单品促销单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
     * 返回活动结束时间在 startTime<=活动结束时间<endTime 之间的所有单品促销单据
     */
*/
type AlibabaWdkItemChangepriceQueryRequest struct {
    model.Params
    // 变价的类型  * 查询变价的单据专用
    type   string
    // 开始时间
    startTime   string
    // 结束时间，结束时间-开始时间不能超过48小时
    endTime   string
    // 渠道店id
    shopId   string
}

// 初始化AlibabaWdkItemChangepriceQueryRequest对象
func NewAlibabaWdkItemChangepriceQueryRequest() *AlibabaWdkItemChangepriceQueryRequest{
    return &AlibabaWdkItemChangepriceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemChangepriceQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.changeprice.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemChangepriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 变价的类型  * 查询变价的单据专用
func (r *AlibabaWdkItemChangepriceQueryRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaWdkItemChangepriceQueryRequest) GetType() string {
    return r.type
}
// StartTime Setter
// 开始时间
func (r *AlibabaWdkItemChangepriceQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaWdkItemChangepriceQueryRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，结束时间-开始时间不能超过48小时
func (r *AlibabaWdkItemChangepriceQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWdkItemChangepriceQueryRequest) GetEndTime() string {
    return r.endTime
}
// ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemChangepriceQueryRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaWdkItemChangepriceQueryRequest) GetShopId() string {
    return r.shopId
}
