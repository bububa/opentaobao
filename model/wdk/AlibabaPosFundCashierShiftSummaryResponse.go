package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
收银换班数据同步 APIResponse
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。
*/
type AlibabaPosFundCashierShiftSummaryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaPosFundCashierShiftSummaryResponse `json:"alibaba_pos_fund_cashier_shift_summary_response,omitempty"` 
    AlibabaPosFundCashierShiftSummaryResponse
}

/* model for simplify = false
type AlibabaPosFundCashierShiftSummaryResponse struct {

    // 接口返回
    
    Result  *struct {
        AlibabaPosFundCashierShiftSummaryResult  *AlibabaPosFundCashierShiftSummaryResult `json:"alibaba_pos_fund_cashier_shift_summary_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaPosFundCashierShiftSummaryResponse struct {

    // 接口返回
    Result   *AlibabaPosFundCashierShiftSummaryResult `json:"result,omitempty"`

}
