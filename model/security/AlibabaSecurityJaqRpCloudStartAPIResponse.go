package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudStartAPIResponse 实人认证云开始认证 API返回值
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
type AlibabaSecurityJaqRpCloudStartAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudStartAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudStartAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudStartAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudStartAPIResponseModel is 实人认证云开始认证 成功返回结果
type AlibabaSecurityJaqRpCloudStartAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_start_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpStartResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudStartAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpCloudStartAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudStartAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudStartAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudStartAPIResponse
func GetAlibabaSecurityJaqRpCloudStartAPIResponse() *AlibabaSecurityJaqRpCloudStartAPIResponse {
	return poolAlibabaSecurityJaqRpCloudStartAPIResponse.Get().(*AlibabaSecurityJaqRpCloudStartAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudStartAPIResponse 将 AlibabaSecurityJaqRpCloudStartAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudStartAPIResponse(v *AlibabaSecurityJaqRpCloudStartAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudStartAPIResponse.Put(v)
}
