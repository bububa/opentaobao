package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务日志打点 APIResponse
alibaba.security.jaq.rp.cloud.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpCloudRphitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudRphitResponse `json:"alibaba_security_jaq_rp_cloud_rphit_response,omitempty"` 
    AlibabaSecurityJaqRpCloudRphitResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudRphitResponse struct {

    // result
    
    Data   string `json:"data,omitempty"`
    

    // result
    
    Data   string `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudRphitResponse struct {

    // result
    Data   string `json:"data,omitempty"`

    // result
    Data   string `json:"data,omitempty"`

}
