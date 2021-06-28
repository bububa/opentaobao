package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向报告 APIResponse
alibaba.scbp.ad.report.get.target.report

定向报告
*/
type AlibabaScbpAdReportGetTargetReportAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdReportGetTargetReportResponse `json:"alibaba_scbp_ad_report_get_target_report_response,omitempty"` 
    AlibabaScbpAdReportGetTargetReportResponse
}

/* model for simplify = false
type AlibabaScbpAdReportGetTargetReportResponse struct {

    // 返回数据
    
    Result  *struct {
        TargetReportDto  *TargetReportDto `json:"target_report_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdReportGetTargetReportResponse struct {

    // 返回数据
    Result   *TargetReportDto `json:"result,omitempty"`

}
