package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证提交认证接口 APIResponse
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpSubmitAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqRpSubmitResponse `json:"alibaba_security_jaq_rp_submit_response,omitempty"`
}

type AlibabaSecurityJaqRpSubmitResponse struct {

    // 结果信息
    Data   *RpSubmitResult `json:"data,omitempty"`

}
