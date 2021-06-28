package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销查询接口 APIResponse
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
type AlibabaWdkPosTradeQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkPosTradeQueryResponse `json:"alibaba_wdk_pos_trade_query_response,omitempty"` 
    AlibabaWdkPosTradeQueryResponse
}

/* model for simplify = false
type AlibabaWdkPosTradeQueryResponse struct {

    // 查询返回结果
    
    Result  *struct {
        FastBuyPosQueryResult  *FastBuyPosQueryResult `json:"fast_buy_pos_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkPosTradeQueryResponse struct {

    // 查询返回结果
    Result   *FastBuyPosQueryResult `json:"result,omitempty"`

}
