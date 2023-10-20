package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiauditresultAPIResponse 乐税审核结果通知 API返回值
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
type AlibabaalihealthdoctorleshuiauditresultAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdoctorleshuiauditresultAPIResponseModel
}

// AlibabaalihealthdoctorleshuiauditresultAPIResponseModel is 乐税审核结果通知 成功返回结果
type AlibabaalihealthdoctorleshuiauditresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_audit_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
