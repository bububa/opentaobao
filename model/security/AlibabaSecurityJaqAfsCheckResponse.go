package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈二次验证接口 APIResponse
alibaba.security.jaq.afs.check

反欺诈二次验证接口
*/
type AlibabaSecurityJaqAfsCheckAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqAfsCheckResponse `json:"alibaba_security_jaq_afs_check_response,omitempty"`
}

type AlibabaSecurityJaqAfsCheckResponse struct {

    // 验证结果
    Data   bool `json:"data,omitempty"`

}
