package foodscan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第二只脚生成报告接口 API返回值 
alibaba.footscan.mini.report.fragment.second

第二只脚生成报告接口
*/
type AlibabaFootscanMiniReportFragmentSecondAPIResponse struct {
    model.CommonResponse
    AlibabaFootscanMiniReportFragmentSecondAPIResponseModel
}

// 第二只脚生成报告接口 成功返回结果
type AlibabaFootscanMiniReportFragmentSecondAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_footscan_mini_report_fragment_second_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 服务出参
    Result   *AlibabaFootscanMiniReportFragmentSecondMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
