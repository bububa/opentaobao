package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaDoctorStatusSaveAPIResponse ca认证获取医师认证结果 API返回值
// alibaba.alihealth.rx.ca.doctor.status.save
//
// ca认证获取医师认证结果
type AlibabaAlihealthRxCaDoctorStatusSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthRxCaDoctorStatusSaveAPIResponseModel
}

// AlibabaAlihealthRxCaDoctorStatusSaveAPIResponseModel is ca认证获取医师认证结果 成功返回结果
type AlibabaAlihealthRxCaDoctorStatusSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_rx_ca_doctor_status_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 错误信息
	ExceptionMessage string `json:"exception_message,omitempty" xml:"exception_message,omitempty"`
	// 出参
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
