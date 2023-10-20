package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudEventAPIResponse 事件上报 API返回值
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
type AlibabaSecurityJaqRpCloudEventAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqRpCloudEventAPIResponseModel
}

// AlibabaSecurityJaqRpCloudEventAPIResponseModel is 事件上报 成功返回结果
type AlibabaSecurityJaqRpCloudEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RpEventResult `json:"result,omitempty" xml:"result,omitempty"`
}
