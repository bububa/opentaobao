package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpSubmitAPIResponse
聚安全实人认证提交认证接口 API返回值
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口 */
type AlibabaSecurityJaqRpSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpSubmitAPIResponseModel
}

// AlibabaSecurityJaqRpSubmitAPIResponseModel is 聚安全实人认证提交认证接口 成功返回结果
type AlibabaSecurityJaqRpSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data *RpSubmitResult `json:"data,omitempty" xml:"data,omitempty"`
}
