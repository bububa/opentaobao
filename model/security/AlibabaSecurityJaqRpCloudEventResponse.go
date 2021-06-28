package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
事件上报 APIResponse
alibaba.security.jaq.rp.cloud.event

事件上报接口
*/
type AlibabaSecurityJaqRpCloudEventAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudEventResponse `json:"alibaba_security_jaq_rp_cloud_event_response,omitempty"` 
    AlibabaSecurityJaqRpCloudEventResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudEventResponse struct {

    // result
    
    Result  *struct {
        RpEventResult  *RpEventResult `json:"rp_event_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudEventResponse struct {

    // result
    Result   *RpEventResult `json:"result,omitempty"`

}
