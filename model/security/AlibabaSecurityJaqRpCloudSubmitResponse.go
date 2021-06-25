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
    Response *AlibabaSecurityJaqRpCloudSubmitResponse `json:"alibaba_security_jaq_rp_cloud_submit_response,omitempty"`
}

type AlibabaSecurityJaqRpCloudSubmitResponse struct {

    // result
    Data   *RpSubmitResult `json:"data,omitempty"`

}
