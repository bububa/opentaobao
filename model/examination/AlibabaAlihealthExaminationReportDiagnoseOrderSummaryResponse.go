package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检报告人工解读总结回传 APIResponse
alibaba.alihealth.examination.report.diagnose.order.summary

记录体检报告人工解读总结
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReportDiagnoseOrderSummaryResponse
}

type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_summary_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
