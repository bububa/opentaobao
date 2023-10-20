package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqOcrImageSyncDetectAPIResponse 聚安全图文识别同步检测接口 API返回值
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
type AlibabaSecurityJaqOcrImageSyncDetectAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqOcrImageSyncDetectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqOcrImageSyncDetectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqOcrImageSyncDetectAPIResponseModel).Reset()
}

// AlibabaSecurityJaqOcrImageSyncDetectAPIResponseModel is 聚安全图文识别同步检测接口 成功返回结果
type AlibabaSecurityJaqOcrImageSyncDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_ocr_image_sync_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqOcrImageDetectResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqOcrImageSyncDetectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqOcrImageSyncDetectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqOcrImageSyncDetectAPIResponse)
	},
}

// GetAlibabaSecurityJaqOcrImageSyncDetectAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqOcrImageSyncDetectAPIResponse
func GetAlibabaSecurityJaqOcrImageSyncDetectAPIResponse() *AlibabaSecurityJaqOcrImageSyncDetectAPIResponse {
	return poolAlibabaSecurityJaqOcrImageSyncDetectAPIResponse.Get().(*AlibabaSecurityJaqOcrImageSyncDetectAPIResponse)
}

// ReleaseAlibabaSecurityJaqOcrImageSyncDetectAPIResponse 将 AlibabaSecurityJaqOcrImageSyncDetectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqOcrImageSyncDetectAPIResponse(v *AlibabaSecurityJaqOcrImageSyncDetectAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqOcrImageSyncDetectAPIResponse.Put(v)
}
