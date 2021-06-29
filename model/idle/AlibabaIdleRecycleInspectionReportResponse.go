package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定报告 APIResponse
alibaba.idle.recycle.inspection.report

回收商鉴定报告
*/
type AlibabaIdleRecycleInspectionReportAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleInspectionReportResponse
}

type AlibabaIdleRecycleInspectionReportResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_inspection_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RecycleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
