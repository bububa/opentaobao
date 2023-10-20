package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpStatusAPIResponse 聚安全实人认证查询状态接口 API返回值
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
type AlibabaSecurityJaqRpStatusAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpStatusAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpStatusAPIResponseModel is 聚安全实人认证查询状态接口 成功返回结果
type AlibabaSecurityJaqRpStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态信息
	Data *RpStatusResultBo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpStatusAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpStatusAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpStatusAPIResponse
func GetAlibabaSecurityJaqRpStatusAPIResponse() *AlibabaSecurityJaqRpStatusAPIResponse {
	return poolAlibabaSecurityJaqRpStatusAPIResponse.Get().(*AlibabaSecurityJaqRpStatusAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpStatusAPIResponse 将 AlibabaSecurityJaqRpStatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpStatusAPIResponse(v *AlibabaSecurityJaqRpStatusAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpStatusAPIResponse.Put(v)
}
