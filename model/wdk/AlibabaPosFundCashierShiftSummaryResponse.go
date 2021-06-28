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
    AlibabaPosFundCashierShiftSummaryResponse
}

type AlibabaPosFundCashierShiftSummaryResponse struct {
    XMLName xml.Name `xml:"alibaba_pos_fund_cashier_shift_summary_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回
    
    Result   *AlibabaPosFundCashierShiftSummaryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
