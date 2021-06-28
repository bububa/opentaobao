package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务提交接口 APIResponse
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpCloudSubmitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudSubmitResponse `json:"alibaba_security_jaq_rp_cloud_submit_response,omitempty"` 
    AlibabaSecurityJaqRpCloudSubmitResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudSubmitResponse struct {

    // result
    
    Data  *struct {
        RpSubmitResult  *RpSubmitResult `json:"rp_submit_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudSubmitResponse struct {

    // result
    Data   *RpSubmitResult `json:"data,omitempty"`

}
