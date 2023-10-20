package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponse 导医通报告解读临时消息接收 API返回值
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
type AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponseModel
}

// AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponseModel is 导医通报告解读临时消息接收 成功返回结果
type AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_tempmessage_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
