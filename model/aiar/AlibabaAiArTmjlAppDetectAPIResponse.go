package aiar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiArTmjlAppDetectAPIResponse 天猫精灵扫一扫入口的服务 API返回值
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
type AlibabaAiArTmjlAppDetectAPIResponse struct {
	model.CommonResponse
	AlibabaAiArTmjlAppDetectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiArTmjlAppDetectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiArTmjlAppDetectAPIResponseModel).Reset()
}

// AlibabaAiArTmjlAppDetectAPIResponseModel is 天猫精灵扫一扫入口的服务 成功返回结果
type AlibabaAiArTmjlAppDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_ar_tmjl_app_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Results string `json:"results,omitempty" xml:"results,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiArTmjlAppDetectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = ""
}

var poolAlibabaAiArTmjlAppDetectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiArTmjlAppDetectAPIResponse)
	},
}

// GetAlibabaAiArTmjlAppDetectAPIResponse 从 sync.Pool 获取 AlibabaAiArTmjlAppDetectAPIResponse
func GetAlibabaAiArTmjlAppDetectAPIResponse() *AlibabaAiArTmjlAppDetectAPIResponse {
	return poolAlibabaAiArTmjlAppDetectAPIResponse.Get().(*AlibabaAiArTmjlAppDetectAPIResponse)
}

// ReleaseAlibabaAiArTmjlAppDetectAPIResponse 将 AlibabaAiArTmjlAppDetectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiArTmjlAppDetectAPIResponse(v *AlibabaAiArTmjlAppDetectAPIResponse) {
	v.Reset()
	poolAlibabaAiArTmjlAppDetectAPIResponse.Put(v)
}
