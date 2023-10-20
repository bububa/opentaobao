package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpStartAPIResponse 聚安全实人认证开始 API返回值
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
type AlibabaSecurityJaqRpStartAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpStartAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpStartAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpStartAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpStartAPIResponseModel is 聚安全实人认证开始 成功返回结果
type AlibabaSecurityJaqRpStartAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_start_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果信息
	Data *RpStartResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpStartAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpStartAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpStartAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpStartAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpStartAPIResponse
func GetAlibabaSecurityJaqRpStartAPIResponse() *AlibabaSecurityJaqRpStartAPIResponse {
	return poolAlibabaSecurityJaqRpStartAPIResponse.Get().(*AlibabaSecurityJaqRpStartAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpStartAPIResponse 将 AlibabaSecurityJaqRpStartAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpStartAPIResponse(v *AlibabaSecurityJaqRpStartAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpStartAPIResponse.Put(v)
}
