package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse
体检报告人工解读总结回传 API返回值
alibaba.alihealth.examination.report.diagnose.order.summary

记录体检报告人工解读总结 */
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel
}

// AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel is 体检报告人工解读总结回传 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
