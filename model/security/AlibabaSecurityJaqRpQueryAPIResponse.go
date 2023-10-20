package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpQueryAPIResponse 聚安全实人认证查询认证结果 API返回值
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
type AlibabaSecurityJaqRpQueryAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpQueryAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpQueryAPIResponseModel is 聚安全实人认证查询认证结果 成功返回结果
type AlibabaSecurityJaqRpQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果信息
	Data *RpAuditResultBo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaSecurityJaqRpQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpQueryAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpQueryAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpQueryAPIResponse
func GetAlibabaSecurityJaqRpQueryAPIResponse() *AlibabaSecurityJaqRpQueryAPIResponse {
	return poolAlibabaSecurityJaqRpQueryAPIResponse.Get().(*AlibabaSecurityJaqRpQueryAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpQueryAPIResponse 将 AlibabaSecurityJaqRpQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpQueryAPIResponse(v *AlibabaSecurityJaqRpQueryAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpQueryAPIResponse.Put(v)
}
