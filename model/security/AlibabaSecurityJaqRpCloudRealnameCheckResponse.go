package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
验证姓名和证件号 APIResponse
alibaba.security.jaq.rp.cloud.realname.check

验证姓名和证件号
*/
type AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudRealnameCheckResponse `json:"alibaba_security_jaq_rp_cloud_realname_check_response,omitempty"` 
    AlibabaSecurityJaqRpCloudRealnameCheckResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudRealnameCheckResponse struct {

    // result
    
    Data  *struct {
        RealNameResult  *RealNameResult `json:"real_name_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudRealnameCheckResponse struct {

    // result
    Data   *RealNameResult `json:"data,omitempty"`

}
