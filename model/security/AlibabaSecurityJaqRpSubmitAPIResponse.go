package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpSubmitAPIResponse 聚安全实人认证提交认证接口 API返回值
// alibaba.security.jaq.rp.submit
//
// 聚安全实人认证提交认证接口
type AlibabaSecurityJaqRpSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpSubmitAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpSubmitAPIResponseModel is 聚安全实人认证提交认证接口 成功返回结果
type AlibabaSecurityJaqRpSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data *RpSubmitResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpSubmitAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpSubmitAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpSubmitAPIResponse
func GetAlibabaSecurityJaqRpSubmitAPIResponse() *AlibabaSecurityJaqRpSubmitAPIResponse {
	return poolAlibabaSecurityJaqRpSubmitAPIResponse.Get().(*AlibabaSecurityJaqRpSubmitAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpSubmitAPIResponse 将 AlibabaSecurityJaqRpSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpSubmitAPIResponse(v *AlibabaSecurityJaqRpSubmitAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpSubmitAPIResponse.Put(v)
}
