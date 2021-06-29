package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商主动通知体检报告 APIResponse
alibaba.alihealth.examination.reserve.report.nofify

服务商主动回传用户的体检报告数据
*/
type AlibabaAlihealthExaminationReserveReportNofifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveReportNofifyResponse
}

type AlibabaAlihealthExaminationReserveReportNofifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_report_nofify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 11
    
    EagleEyeTraceId   string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`

    
    // 返回数据对象
    
    Data   *ReserveReportResponse `json:"data,omitempty" xml:"data,omitempty"`

    
}
