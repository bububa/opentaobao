package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse 异步开方处方查询 API返回值
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel is 异步开方处方查询 成功返回结果
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_asyncprescribe_prescription_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse)
	},
}

// GetAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse
func GetAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse() *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse {
	return poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse.Get().(*AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse)
}

// ReleaseAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse 将 AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse(v *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse.Put(v)
}
