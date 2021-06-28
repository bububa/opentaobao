package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资金合规商家账单 APIResponse
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单
*/
type AlibabaWdkOrderFinanceBillQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderFinanceBillQueryResponse `json:"alibaba_wdk_order_finance_bill_query_response,omitempty"` 
    AlibabaWdkOrderFinanceBillQueryResponse
}

/* model for simplify = false
type AlibabaWdkOrderFinanceBillQueryResponse struct {

    // 出参
    
    Result  *struct {
        WdkOpenOrderFinanceBillQueryResult  *WdkOpenOrderFinanceBillQueryResult `json:"wdk_open_order_finance_bill_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderFinanceBillQueryResponse struct {

    // 出参
    Result   *WdkOpenOrderFinanceBillQueryResult `json:"result,omitempty"`

}
