package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
事件上报 APIResponse
alibaba.security.jaq.rp.cloud.event

事件上报接口
*/
type AlibabaSecurityJaqRpCloudEventAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpCloudEventResponse
}

type AlibabaSecurityJaqRpCloudEventResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_event_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *RpEventResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
