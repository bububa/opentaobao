package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按照价格变更时间段，查询会变更价格的单据的商品 APIResponse
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
type AlibabaWdkItemChangepriceQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkItemChangepriceQueryResponse `json:"alibaba_wdk_item_changeprice_query_response,omitempty"` 
    AlibabaWdkItemChangepriceQueryResponse
}

/* model for simplify = false
type AlibabaWdkItemChangepriceQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaWdkItemChangepriceQueryResult  *AlibabaWdkItemChangepriceQueryResult `json:"alibaba_wdk_item_changeprice_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkItemChangepriceQueryResponse struct {

    // 接口返回model
    Result   *AlibabaWdkItemChangepriceQueryResult `json:"result,omitempty"`

}
