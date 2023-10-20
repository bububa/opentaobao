package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpGetverifytokenAPIResponse 聚安全实人认证获取认证会话token API返回值
// alibaba.security.jaq.rp.getverifytoken
//
// 聚安全实人认证获取认证会话token
type AlibabaSecurityJaqRpGetverifytokenAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpGetverifytokenAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpGetverifytokenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpGetverifytokenAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpGetverifytokenAPIResponseModel is 聚安全实人认证获取认证会话token 成功返回结果
type AlibabaSecurityJaqRpGetverifytokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_getverifytoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// token信息
	Data *RpInitResultBo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpGetverifytokenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpGetverifytokenAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpGetverifytokenAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpGetverifytokenAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpGetverifytokenAPIResponse
func GetAlibabaSecurityJaqRpGetverifytokenAPIResponse() *AlibabaSecurityJaqRpGetverifytokenAPIResponse {
	return poolAlibabaSecurityJaqRpGetverifytokenAPIResponse.Get().(*AlibabaSecurityJaqRpGetverifytokenAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpGetverifytokenAPIResponse 将 AlibabaSecurityJaqRpGetverifytokenAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpGetverifytokenAPIResponse(v *AlibabaSecurityJaqRpGetverifytokenAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpGetverifytokenAPIResponse.Put(v)
}
