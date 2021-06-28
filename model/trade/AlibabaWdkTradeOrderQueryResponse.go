package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询外部交易订单接口 APIResponse
alibaba.wdk.trade.order.query

通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTradeOrderQueryResponse `json:"alibaba_wdk_trade_order_query_response,omitempty"` 
    AlibabaWdkTradeOrderQueryResponse
}

/* model for simplify = false
type AlibabaWdkTradeOrderQueryResponse struct {

    // 查询结果
    
    Result  *struct {
        TradeOrderQueryResult  *TradeOrderQueryResult `json:"trade_order_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkTradeOrderQueryResponse struct {

    // 查询结果
    Result   *TradeOrderQueryResult `json:"result,omitempty"`

}
