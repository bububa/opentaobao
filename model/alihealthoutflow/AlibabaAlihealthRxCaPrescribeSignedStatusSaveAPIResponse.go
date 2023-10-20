package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponse 处方ca认证 API返回值
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
type AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponseModel
}

// AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponseModel is 处方ca认证 成功返回结果
type AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_rx_ca_prescribe_signed_status_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 错误信息
	ExceptionMessage string `json:"exception_message,omitempty" xml:"exception_message,omitempty"`
	// 出参
	Data *YwxCommonVo `json:"data,omitempty" xml:"data,omitempty"`
}
