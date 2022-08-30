package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponse 处方ca认证-厂商通知接口 API返回值
// alibaba.alihealth.rx.ca.device.sign.status.save
//
// 处方ca认证-厂商通知接口
type AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponseModel
}

// AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponseModel is 处方ca认证-厂商通知接口 成功返回结果
type AlibabaAlihealthRxCaDeviceSignStatusSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_rx_ca_device_sign_status_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 错误信息
	ExceptionMessage string `json:"exception_message,omitempty" xml:"exception_message,omitempty"`
	// 出参
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
