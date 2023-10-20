package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPatientinfoSyncAPIResponse 处方外流-患者基础信息同步 API返回值
// alibaba.alihealth.outflow.patientinfo.sync
//
// 阿里健康-处方外流-对外提供同步患者基础信息功能
type AlibabaAlihealthOutflowPatientinfoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPatientinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel is 处方外流-患者基础信息同步 成功返回结果
type AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_patientinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowPatientinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowPatientinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowPatientinfoSyncAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowPatientinfoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowPatientinfoSyncAPIResponse
func GetAlibabaAlihealthOutflowPatientinfoSyncAPIResponse() *AlibabaAlihealthOutflowPatientinfoSyncAPIResponse {
	return poolAlibabaAlihealthOutflowPatientinfoSyncAPIResponse.Get().(*AlibabaAlihealthOutflowPatientinfoSyncAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowPatientinfoSyncAPIResponse 将 AlibabaAlihealthOutflowPatientinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowPatientinfoSyncAPIResponse(v *AlibabaAlihealthOutflowPatientinfoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowPatientinfoSyncAPIResponse.Put(v)
}
