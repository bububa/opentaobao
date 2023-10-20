package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiticketvalidAPIResponse 乐税token验证 API返回值
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
type AlibabaalihealthdoctorleshuiticketvalidAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdoctorleshuiticketvalidAPIResponseModel
}

// AlibabaalihealthdoctorleshuiticketvalidAPIResponseModel is 乐税token验证 成功返回结果
type AlibabaalihealthdoctorleshuiticketvalidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_ticket_valid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
