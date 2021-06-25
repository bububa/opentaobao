package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品当前价格 APIResponse
alibaba.wdk.item.currentprice.query

通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量<=20,返回结果key为skuCode
*/
type AlibabaWdkItemCurrentpriceQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkItemCurrentpriceQueryResponse `json:"alibaba_wdk_item_currentprice_query_response,omitempty"`
}

type AlibabaWdkItemCurrentpriceQueryResponse struct {

    // 接口返回model
    Result   *AlibabaWdkItemCurrentpriceQueryResult `json:"result,omitempty"`

}
