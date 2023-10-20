package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse 导医通报告解读临时消息接收 API返回值
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponseModel is 导医通报告解读临时消息接收 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_tempmessage_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse.Put(v)
}
