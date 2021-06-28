package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
收银换班数据同步 APIResponse
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。
*/
type AlibabaPosFundCashierShiftSummaryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_pos_fund_cashier_shift_summary_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回
    
    Result   *AlibabaPosFundCashierShiftSummaryResult `json:"result,omitempty" xml:"