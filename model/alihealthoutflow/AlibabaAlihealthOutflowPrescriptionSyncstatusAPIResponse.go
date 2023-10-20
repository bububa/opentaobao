package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse 处方外流-处方状态同步 API返回值
// alibaba.alihealth.outflow.prescription.syncstatus
//
// 阿里健康-处方外流-对外提供同步处方状态功能
type AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponseModel is 处方外流-处方状态同步 成功返回结果
type AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_syncstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse
func GetAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse() *AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse {
	return poolAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse.Get().(*AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse 将 AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse(v *AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse.Put(v)
}
