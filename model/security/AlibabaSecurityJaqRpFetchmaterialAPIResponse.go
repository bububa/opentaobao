package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpFetchmaterialAPIResponse 聚安全实人认证获取结果接口 API返回值
// alibaba.security.jaq.rp.fetchmaterial
//
// 聚安全实人认证获取结果接口
type AlibabaSecurityJaqRpFetchmaterialAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpFetchmaterialAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpFetchmaterialAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpFetchmaterialAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpFetchmaterialAPIResponseModel is 聚安全实人认证获取结果接口 成功返回结果
type AlibabaSecurityJaqRpFetchmaterialAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_fetchmaterial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpFetchmaterialAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolAlibabaSecurityJaqRpFetchmaterialAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpFetchmaterialAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpFetchmaterialAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpFetchmaterialAPIResponse
func GetAlibabaSecurityJaqRpFetchmaterialAPIResponse() *AlibabaSecurityJaqRpFetchmaterialAPIResponse {
	return poolAlibabaSecurityJaqRpFetchmaterialAPIResponse.Get().(*AlibabaSecurityJaqRpFetchmaterialAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpFetchmaterialAPIResponse 将 AlibabaSecurityJaqRpFetchmaterialAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpFetchmaterialAPIResponse(v *AlibabaSecurityJaqRpFetchmaterialAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpFetchmaterialAPIResponse.Put(v)
}
