package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponse 获取报告解读url API返回值
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
type AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponseModel
}

// AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponseModel is 获取报告解读url 成功返回结果
type AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_diagnoseurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信中跳转url
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
