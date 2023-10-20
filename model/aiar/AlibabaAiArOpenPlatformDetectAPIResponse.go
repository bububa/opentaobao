package aiar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiArOpenPlatformDetectAPIResponse AR开发者平台marker图片检测服务 API返回值
// alibaba.ai.ar.open.platform.detect
//
// AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用
type AlibabaAiArOpenPlatformDetectAPIResponse struct {
	model.CommonResponse
	AlibabaAiArOpenPlatformDetectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiArOpenPlatformDetectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiArOpenPlatformDetectAPIResponseModel).Reset()
}

// AlibabaAiArOpenPlatformDetectAPIResponseModel is AR开发者平台marker图片检测服务 成功返回结果
type AlibabaAiArOpenPlatformDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_ar_open_platform_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiArOpenPlatformDetectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaAiArOpenPlatformDetectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiArOpenPlatformDetectAPIResponse)
	},
}

// GetAlibabaAiArOpenPlatformDetectAPIResponse 从 sync.Pool 获取 AlibabaAiArOpenPlatformDetectAPIResponse
func GetAlibabaAiArOpenPlatformDetectAPIResponse() *AlibabaAiArOpenPlatformDetectAPIResponse {
	return poolAlibabaAiArOpenPlatformDetectAPIResponse.Get().(*AlibabaAiArOpenPlatformDetectAPIResponse)
}

// ReleaseAlibabaAiArOpenPlatformDetectAPIResponse 将 AlibabaAiArOpenPlatformDetectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiArOpenPlatformDetectAPIResponse(v *AlibabaAiArOpenPlatformDetectAPIResponse) {
	v.Reset()
	poolAlibabaAiArOpenPlatformDetectAPIResponse.Put(v)
}
