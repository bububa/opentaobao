package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse 异步开方处方详情 API返回值
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
type AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponseModel).Reset()
}

// AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponseModel is 异步开方处方详情 成功返回结果
type AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_asyncprescribe_prescription_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse)
	},
}

// GetAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse
func GetAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse() *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse {
	return poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse.Get().(*AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse)
}

// ReleaseAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse 将 AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse(v *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse.Put(v)
}
