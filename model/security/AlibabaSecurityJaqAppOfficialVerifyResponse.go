package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全验证官方应用接口 APIResponse
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用
*/
type AlibabaSecurityJaqAppOfficialVerifyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqAppOfficialVerifyResponse `json:"alibaba_security_jaq_app_official_verify_response,omitempty"` 
    AlibabaSecurityJaqAppOfficialVerifyResponse
}

/* model for simplify = false
type AlibabaSecurityJaqAppOfficialVerifyResponse struct {

    // result
    
    Result  *struct {
        OfficialAppVerifyResponse  *OfficialAppVerifyResponse `json:"official_app_verify_response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqAppOfficialVerifyResponse struct {

    // result
    Result   *OfficialAppVerifyResponse `json:"result,omitempty"`

}
