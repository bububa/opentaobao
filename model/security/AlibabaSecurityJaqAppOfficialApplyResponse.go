package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全官方应用申请 APIResponse
alibaba.security.jaq.app.official.apply

官方应用申请接口
*/
type AlibabaSecurityJaqAppOfficialApplyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqAppOfficialApplyResponse `json:"alibaba_security_jaq_app_official_apply_response,omitempty"`
}

type AlibabaSecurityJaqAppOfficialApplyResponse struct {

    // 申请结果
    Result   *OfficialAppApplyResponse `json:"result,omitempty"`

}
