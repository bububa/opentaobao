package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云开始认证 APIResponse
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpCloudStartAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqRpCloudStartResponse `json:"alibaba_security_jaq_rp_cloud_start_response,omitempty"`
}

type AlibabaSecurityJaqRpCloudStartResponse struct {

    // result
    Data   *RpStartResult `json:"data,omitempty"`

}
