package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页拉取订单数据 APIResponse
alibaba.wdk.trade.order.balance.bill.query

提供接口供外部调用，分页拉取订单数据
*/
type AlibabaWdkTradeOrderBalanceBillQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_trade_order_balance_bill_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ApiResult
    
    ApiResult   *AlibabaWdkTradeOrderBalanceBillQueryApiResult `json:"api_result,omitempty" xml:"