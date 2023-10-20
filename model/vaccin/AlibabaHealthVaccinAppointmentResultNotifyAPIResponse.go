package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinAppointmentResultNotifyAPIResponse 通知预约结果 API返回值
// alibaba.health.vaccin.appointment.result.notify
//
// 和ISV合作，需ISV回传预约结果。
type AlibabaHealthVaccinAppointmentResultNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinAppointmentResultNotifyAPIResponseModel
}

// AlibabaHealthVaccinAppointmentResultNotifyAPIResponseModel is 通知预约结果 成功返回结果
type AlibabaHealthVaccinAppointmentResultNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_appointment_result_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 1
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 1
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
