package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
账户报告 APIResponse
alibaba.scbp.ad.report.get.account.report

账户报告
*/
type AlibabaScbpAdReportGetAccountReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_report_get_account_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回参数
    
    Result   *AccountReportDto `json:"result,omitempty" xml:"