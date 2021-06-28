package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
账户报告 APIResponse
alibaba.scbp.ad.report.get.account.report

账户报告
*/
type AlibabaScbpAdReportGetAccountReportAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdReportGetAccountReportResponse `json:"alibaba_scbp_ad_report_get_account_report_response,omitempty"` 
    AlibabaScbpAdReportGetAccountReportResponse
}

/* model for simplify = false
type AlibabaScbpAdReportGetAccountReportResponse struct {

    // 返回参数
    
    Result  *struct {
        AccountReportDto  *AccountReportDto `json:"account_report_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdReportGetAccountReportResponse struct {

    // 返回参数
    Result   *AccountReportDto `json:"result,omitempty"`

}
