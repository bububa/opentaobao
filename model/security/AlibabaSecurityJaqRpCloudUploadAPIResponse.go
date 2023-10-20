package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudUploadAPIResponse 实人认证云上传接口 API返回值
// alibaba.security.jaq.rp.cloud.upload
//
// 聚安全实人认证上传认证信息
type AlibabaSecurityJaqRpCloudUploadAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudUploadAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudUploadAPIResponseModel is 实人认证云上传接口 成功返回结果
type AlibabaSecurityJaqRpCloudUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpUploadResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpCloudUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudUploadAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudUploadAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudUploadAPIResponse
func GetAlibabaSecurityJaqRpCloudUploadAPIResponse() *AlibabaSecurityJaqRpCloudUploadAPIResponse {
	return poolAlibabaSecurityJaqRpCloudUploadAPIResponse.Get().(*AlibabaSecurityJaqRpCloudUploadAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudUploadAPIResponse 将 AlibabaSecurityJaqRpCloudUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudUploadAPIResponse(v *AlibabaSecurityJaqRpCloudUploadAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudUploadAPIResponse.Put(v)
}
