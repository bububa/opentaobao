package security

import (
	"encoding/xml"

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

// AlibabaSecurityJaqRpRphitAPIResponseModel is 聚安全-实人认证日志打点接口 成功返回结果
type AlibabaSecurityJaqRpRphitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_rphit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
