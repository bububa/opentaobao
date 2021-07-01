package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpFetchmaterialAPIResponse
聚安全实人认证获取结果接口 API返回值
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口 */
type AlibabaSecurityJaqRpFetchmaterialAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpFetchmaterialAPIResponseModel
}

// AlibabaSecurityJaqRpFetchmaterialAPIResponseModel is 聚安全实人认证获取结果接口 成功返回结果
type AlibabaSecurityJaqRpFetchmaterialAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_fetchmaterial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
