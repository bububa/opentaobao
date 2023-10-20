package alihealthoutflow

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel is 处方外流-创建处方 成功返回结果
type AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowPrescriptionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowPrescriptionCreateAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowPrescriptionCreateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionCreateAPIResponse
func GetAlibabaAlihealthOutflowPrescriptionCreateAPIResponse() *AlibabaAlihealthOutflowPrescriptionCreateAPIResponse {
	return poolAlibabaAlihealthOutflowPrescriptionCreateAPIResponse.Get().(*AlibabaAlihealthOutflowPrescriptionCreateAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionCreateAPIResponse 将 AlibabaAlihealthOutflowPrescriptionCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionCreateAPIResponse(v *AlibabaAlihealthOutflowPrescriptionCreateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionCreateAPIResponse.Put(v)
}
