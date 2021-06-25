package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销的订单活动信息查询 APIResponse
alibaba.wdk.bm.trade.activity.query

品牌营销的订单活动信息查询
*/
type AlibabaWdkBmTradeActivityQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkBmTradeActivityQueryResponse `json:"alibaba_wdk_bm_trade_activity_query_response,omitempty"`
}

type AlibabaWdkBmTradeActivityQueryResponse struct {

    // 结果数据
    Result   *BmResult `json:"result,omitempty"`

}
