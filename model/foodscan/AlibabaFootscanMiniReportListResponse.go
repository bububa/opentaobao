package foodscan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询报告列表 API返回值 
alibaba.footscan.mini.report.list

查询报告列表
*/
type AlibabaFootscanMiniReportListAPIResponse struct {
    model.CommonResponse
    AlibabaFootscanMiniReportListResponse
}

// 查询报告列表 成功返回结果
type AlibabaFootscanMiniReportListResponse struct {
    XMLName xml.Name `xml:"alibaba_footscan_mini_report_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 服务出参
    Result   *AlibabaFootscanMiniReportListMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
