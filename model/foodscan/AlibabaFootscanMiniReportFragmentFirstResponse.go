package foodscan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第一只脚生成报告接口 APIResponse
alibaba.footscan.mini.report.fragment.first

第一只脚生成报告接口
*/
type AlibabaFootscanMiniReportFragmentFirstAPIResponse struct {
    model.CommonResponse
    AlibabaFootscanMiniReportFragmentFirstResponse
}

type AlibabaFootscanMiniReportFragmentFirstResponse struct {
    XMLName xml.Name `xml:"alibaba_footscan_mini_report_fragment_first_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *AlibabaFootscanMiniReportFragmentFirstMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
