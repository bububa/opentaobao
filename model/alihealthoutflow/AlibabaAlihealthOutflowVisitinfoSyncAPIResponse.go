package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowVisitinfoSyncAPIResponse 处方外流-问诊、处方同步 API返回值
// alibaba.alihealth.outflow.visitinfo.sync
//
// 阿里健康-处方外流-对外提供问诊、处方功能
type AlibabaAlihealthOutflowVisitinfoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowVisitinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel is 处方外流-问诊、处方同步 成功返回结果
type AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_visitinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthOutflowVisitinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthOutflowVisitinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthOutflowVisitinfoSyncAPIResponse)
	},
}

// GetAlibabaAlihealthOutflowVisitinfoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthOutflowVisitinfoSyncAPIResponse
func GetAlibabaAlihealthOutflowVisitinfoSyncAPIResponse() *AlibabaAlihealthOutflowVisitinfoSyncAPIResponse {
	return poolAlibabaAlihealthOutflowVisitinfoSyncAPIResponse.Get().(*AlibabaAlihealthOutflowVisitinfoSyncAPIResponse)
}

// ReleaseAlibabaAlihealthOutflowVisitinfoSyncAPIResponse 将 AlibabaAlihealthOutflowVisitinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthOutflowVisitinfoSyncAPIResponse(v *AlibabaAlihealthOutflowVisitinfoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthOutflowVisitinfoSyncAPIResponse.Put(v)
}
