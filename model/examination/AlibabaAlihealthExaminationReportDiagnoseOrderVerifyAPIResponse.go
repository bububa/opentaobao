package examination

import (
	"encoding/xml"

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

// AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel is 报告解读令牌校验 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
