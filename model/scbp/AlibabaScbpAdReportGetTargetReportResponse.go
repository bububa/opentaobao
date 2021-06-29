package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向报告 API返回值 
alibaba.scbp.ad.report.get.target.report

定向报告
*/
type AlibabaScbpAdReportGetTargetReportAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdReportGetTargetReportResponse
}

// 定向报告 成功返回结果
type AlibabaScbpAdReportGetTargetReportResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_target_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *TargetReportDTO `json:"result,omitempty" xml:"result,omitempty"`
}
