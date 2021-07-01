package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudSubmitAPIResponse
实人认证云服务提交接口 API返回值
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口 */
type AlibabaSecurityJaqRpCloudSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudSubmitAPIResponseModel
}

// AlibabaSecurityJaqRpCloudSubmitAPIResponseModel is 实人认证云服务提交接口 成功返回结果
type AlibabaSecurityJaqRpCloudSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpSubmitResult `json:"data,omitempty" xml:"data,omitempty"`
}
