package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpUploadAPIResponse 聚安全实人认证上传认证信息 API返回值
// alibaba.security.jaq.rp.upload
//
// 聚安全实人认证上传认证信息
type AlibabaSecurityJaqRpUploadAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpUploadAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpUploadAPIResponseModel is 聚安全实人认证上传认证信息 成功返回结果
type AlibabaSecurityJaqRpUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Data *RpUploadResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpUploadAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpUploadAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpUploadAPIResponse
func GetAlibabaSecurityJaqRpUploadAPIResponse() *AlibabaSecurityJaqRpUploadAPIResponse {
	return poolAlibabaSecurityJaqRpUploadAPIResponse.Get().(*AlibabaSecurityJaqRpUploadAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpUploadAPIResponse 将 AlibabaSecurityJaqRpUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpUploadAPIResponse(v *AlibabaSecurityJaqRpUploadAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpUploadAPIResponse.Put(v)
}
