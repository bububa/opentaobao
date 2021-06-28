package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页拉取订单数据 APIResponse
alibaba.wdk.trade.order.balance.bill.query

提供接口供外部调用，分页拉取订单数据
*/
type AlibabaWdkTradeOrderBalanceBillQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTradeOrderBalanceBillQueryResponse `json:"alibaba_wdk_trade_order_balance_bill_query_response,omitempty"` 
    AlibabaWdkTradeOrderBalanceBillQueryResponse
}

/* model for simplify = false
type AlibabaWdkTradeOrderBalanceBillQueryResponse struct {

    // ApiResult
    
    ApiResult  *struct {
        AlibabaWdkTradeOrderBalanceBillQueryApiResult  *AlibabaWdkTradeOrderBalanceBillQueryApiResult `json:"alibaba_wdk_trade_order_balance_bill_query_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkTradeOrderBalanceBillQueryResponse struct {

    // ApiResult
    ApiResult   *AlibabaWdkTradeOrderBalanceBillQueryApiResult `json:"api_result,omitempty"`

}
