package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiapplynotifyAPIResponse 申请单审核结果通知 API返回值
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
type AlibabaalihealthdoctorleshuiapplynotifyAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdoctorleshuiapplynotifyAPIResponseModel
}

// AlibabaalihealthdoctorleshuiapplynotifyAPIResponseModel is 申请单审核结果通知 成功返回结果
type AlibabaalihealthdoctorleshuiapplynotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_apply_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
