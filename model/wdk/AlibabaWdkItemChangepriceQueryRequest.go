package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照价格变更时间段，查询会变更价格的单据的商品 APIRequest
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

func NewAlibabaWdkItemChangepriceQueryRequest() *AlibabaWdkItemChangepriceQueryRequest{
    return &AlibabaWdkItemChangepriceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.changeprice.query"
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemChangepriceQueryRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetType() string {
    return r.type
}

func (r *AlibabaWdkItemChangepriceQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetStartTime() string {
    return r.startTime
}

func (r *AlibabaWdkItemChangepriceQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *AlibabaWdkItemChangepriceQueryRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkItemChangepriceQueryRequest) GetShopId() string {
    return r.shopId
}

