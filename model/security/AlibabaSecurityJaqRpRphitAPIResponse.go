package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpRphitAPIResponse 聚安全-实人认证日志打点接口 API返回值
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
type AlibabaSecurityJaqRpRphitAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpRphitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpRphitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpRphitAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpRphitAPIResponseModel is 聚安全-实人认证日志打点接口 成功返回结果
type AlibabaSecurityJaqRpRphitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_rphit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpRphitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolAlibabaSecurityJaqRpRphitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpRphitAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpRphitAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpRphitAPIResponse
func GetAlibabaSecurityJaqRpRphitAPIResponse() *AlibabaSecurityJaqRpRphitAPIResponse {
	return poolAlibabaSecurityJaqRpRphitAPIResponse.Get().(*AlibabaSecurityJaqRpRphitAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpRphitAPIResponse 将 AlibabaSecurityJaqRpRphitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpRphitAPIResponse(v *AlibabaSecurityJaqRpRphitAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpRphitAPIResponse.Put(v)
}
