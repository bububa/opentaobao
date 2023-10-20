package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse 报告解读订单状态更新 API返回值
// alibaba.alihealth.examination.report.diagnose.order.status
//
// 报告解读订单状态更新
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel is 报告解读订单状态更新 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse.Put(v)
}
