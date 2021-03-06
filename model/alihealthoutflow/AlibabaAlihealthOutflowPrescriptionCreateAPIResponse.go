package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionCreateAPIResponse 处方外流-创建处方 API返回值
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
type AlibabaAlihealthOutflowPrescriptionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel
}

// AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel is 处方外流-创建处方 成功返回结果
type AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
