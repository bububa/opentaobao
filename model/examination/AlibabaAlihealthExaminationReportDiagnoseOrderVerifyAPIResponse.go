package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse 报告解读令牌校验 API返回值
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
type AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel is 报告解读令牌校验 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse.Put(v)
}
