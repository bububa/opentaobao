package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudRphitAPIResponse 实人认证云服务日志打点 API返回值
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
type AlibabaSecurityJaqRpCloudRphitAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudRphitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudRphitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqRpCloudRphitAPIResponseModel).Reset()
}

// AlibabaSecurityJaqRpCloudRphitAPIResponseModel is 实人认证云服务日志打点 成功返回结果
type AlibabaSecurityJaqRpCloudRphitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_rphit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqRpCloudRphitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolAlibabaSecurityJaqRpCloudRphitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqRpCloudRphitAPIResponse)
	},
}

// GetAlibabaSecurityJaqRpCloudRphitAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudRphitAPIResponse
func GetAlibabaSecurityJaqRpCloudRphitAPIResponse() *AlibabaSecurityJaqRpCloudRphitAPIResponse {
	return poolAlibabaSecurityJaqRpCloudRphitAPIResponse.Get().(*AlibabaSecurityJaqRpCloudRphitAPIResponse)
}

// ReleaseAlibabaSecurityJaqRpCloudRphitAPIResponse 将 AlibabaSecurityJaqRpCloudRphitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudRphitAPIResponse(v *AlibabaSecurityJaqRpCloudRphitAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudRphitAPIResponse.Put(v)
}
