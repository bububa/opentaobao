package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservereportnofifyAPIResponse 服务商主动通知体检报告 API返回值
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
type AlibabaalihealthexaminationreservereportnofifyAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreservereportnofifyAPIResponseModel
}

// AlibabaalihealthexaminationreservereportnofifyAPIResponseModel is 服务商主动通知体检报告 成功返回结果
type AlibabaalihealthexaminationreservereportnofifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_report_nofify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	EagleEyeTraceId string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
	// 返回数据对象
	Data *ReserveReportResponse `json:"data,omitempty" xml:"data,omitempty"`
}
