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
    Response *AlibabaWdkOrderFinanceBillQueryResponse `json:"alibaba_wdk_order_finance_bill_query_response,omitempty"`
}

type AlibabaWdkOrderFinanceBillQueryResponse struct {

    // 出参
    Result   *WdkOpenOrderFinanceBillQueryResult `json:"result,omitempty"`

}
