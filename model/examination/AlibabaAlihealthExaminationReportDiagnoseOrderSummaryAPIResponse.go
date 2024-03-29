package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse 体检报告人工解读总结回传 API返回值
// alibaba.alihealth.examination.report.diagnose.order.summary
//
// 记录体检报告人工解读总结
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel is 体检报告人工解读总结回传 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse.Put(v)
}
