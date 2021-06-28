package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户查询加固结果 APIResponse
alibaba.security.jaq.app.shieldresult.get

用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
*/
type AlibabaSecurityJaqAppShieldresultGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqAppShieldresultGetResponse `json:"alibaba_security_jaq_app_shieldresult_get_response,omitempty"` 
    AlibabaSecurityJaqAppShieldresultGetResponse
}

/* model for simplify = false
type AlibabaSecurityJaqAppShieldresultGetResponse struct {

    // 应用加固结果
    
    Result  *struct {
        ShieldResult  *ShieldResult `json:"shield_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqAppShieldresultGetResponse struct {

    // 应用加固结果
    Result   *ShieldResult `json:"result,omitempty"`

}
