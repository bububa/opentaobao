package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导医通报告解读临时消息接收 APIResponse
alibaba.alihealth.examination.report.diagnose.tempmessage.receive

导医通报告解读临时消息接收
*/
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveResponse
}

type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_tempmessage_receive_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
