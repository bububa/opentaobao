package aiar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiArServiceDetectAPIResponse ailab AR图像检索 API返回值
// alibaba.ai.ar.service.detect
//
// ailab AR图像检索
type AlibabaAiArServiceDetectAPIResponse struct {
	model.CommonResponse
	AlibabaAiArServiceDetectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiArServiceDetectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiArServiceDetectAPIResponseModel).Reset()
}

// AlibabaAiArServiceDetectAPIResponseModel is ailab AR图像检索 成功返回结果
type AlibabaAiArServiceDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_ar_service_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Results string `json:"results,omitempty" xml:"results,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiArServiceDetectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = ""
}

var poolAlibabaAiArServiceDetectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiArServiceDetectAPIResponse)
	},
}

// GetAlibabaAiArServiceDetectAPIResponse 从 sync.Pool 获取 AlibabaAiArServiceDetectAPIResponse
func GetAlibabaAiArServiceDetectAPIResponse() *AlibabaAiArServiceDetectAPIResponse {
	return poolAlibabaAiArServiceDetectAPIResponse.Get().(*AlibabaAiArServiceDetectAPIResponse)
}

// ReleaseAlibabaAiArServiceDetectAPIResponse 将 AlibabaAiArServiceDetectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiArServiceDetectAPIResponse(v *AlibabaAiArServiceDetectAPIResponse) {
	v.Reset()
	poolAlibabaAiArServiceDetectAPIResponse.Put(v)
}
