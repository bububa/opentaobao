package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse 聚安全获取异步图文识别结果接口 API返回值
// alibaba.security.jaq.ocr.image.async.detect.results.fetch
//
// 获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponseModel is 聚安全获取异步图文识别结果接口 成功返回结果
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_ocr_image_async_detect_results_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqImageDetectResultCollection `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse)
	},
}

// GetAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse
func GetAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse {
	return poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse.Get().(*AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse)
}

// ReleaseAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse 将 AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse(v *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse.Put(v)
}
