package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqPornImageSyncDetectAPIResponse 聚安全智能鉴黄同步检测接口 API返回值
// alibaba.security.jaq.porn.image.sync.detect
//
// 同步黄图图像检测接口
type AlibabaSecurityJaqPornImageSyncDetectAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqPornImageSyncDetectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel).Reset()
}

// AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel is 聚安全智能鉴黄同步检测接口 成功返回结果
type AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_porn_image_sync_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqPornImageDetectResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqPornImageSyncDetectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqPornImageSyncDetectAPIResponse)
	},
}

// GetAlibabaSecurityJaqPornImageSyncDetectAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqPornImageSyncDetectAPIResponse
func GetAlibabaSecurityJaqPornImageSyncDetectAPIResponse() *AlibabaSecurityJaqPornImageSyncDetectAPIResponse {
	return poolAlibabaSecurityJaqPornImageSyncDetectAPIResponse.Get().(*AlibabaSecurityJaqPornImageSyncDetectAPIResponse)
}

// ReleaseAlibabaSecurityJaqPornImageSyncDetectAPIResponse 将 AlibabaSecurityJaqPornImageSyncDetectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqPornImageSyncDetectAPIResponse(v *AlibabaSecurityJaqPornImageSyncDetectAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqPornImageSyncDetectAPIResponse.Put(v)
}
