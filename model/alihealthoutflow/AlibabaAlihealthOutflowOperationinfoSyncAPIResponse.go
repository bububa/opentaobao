package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowOperationinfoSyncAPIResponse 处方外流-操作信息同步 API返回值
// alibaba.alihealth.outflow.operationinfo.sync
//
// 阿里健康-处方外流-对外提供同步操作信息功能
type AlibabaAlihealthOutflowOperationinfoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowOperationinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowOperationinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowOperationinfoSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowOperationinfoSyncAPIResponseModel is 处方外流-操作信息同步 成功返回结果
type AlibabaAlihealthOutflowOperationinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_operationinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowOperationinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowOperationinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowOperationinfoSyncAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowOperationinfoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowOperationinfoSyncAPIResponse
func GetAlibabaAlihealthOutflowOperationinfoSyncAPIResponse() *AlibabaAlihealthOutflowOperationinfoSyncAPIResponse {
	return poolAlibabaAlihealthOutflowOperationinfoSyncAPIResponse.Get().(*AlibabaAlihealthOutflowOperationinfoSyncAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowOperationinfoSyncAPIResponse 将 AlibabaAlihealthOutflowOperationinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowOperationinfoSyncAPIResponse(v *AlibabaAlihealthOutflowOperationinfoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowOperationinfoSyncAPIResponse.Put(v)
}
