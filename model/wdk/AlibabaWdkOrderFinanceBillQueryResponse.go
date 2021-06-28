package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资金合规商家账单 APIResponse
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单
*/
type AlibabaWdkOrderFinanceBillQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_order_finance_bill_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *WdkOpenOrderFinanceBillQueryResult `json:"result,omitempty" xml:"