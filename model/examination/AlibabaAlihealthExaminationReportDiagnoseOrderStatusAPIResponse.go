package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse
报告解读订单状态更新 API返回值
alibaba.alihealth.examination.report.diagnose.order.status

报告解读订单状态更新 */
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel
}

// AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel is 报告解读订单状态更新 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
