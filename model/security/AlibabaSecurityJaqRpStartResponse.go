package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证开始 APIResponse
alibaba.security.jaq.rp.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpStartAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqRpStartResponse `json:"alibaba_security_jaq_rp_start_response,omitempty"`
}

type AlibabaSecurityJaqRpStartResponse struct {

    // 返回结果信息
    Data   *RpStartResult `json:"data,omitempty"`

}
