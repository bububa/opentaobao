package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取结果接口 APIResponse
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
type AlibabaSecurityJaqRpFetchmaterialAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpFetchmaterialResponse `json:"alibaba_security_jaq_rp_fetchmaterial_response,omitempty"` 
    AlibabaSecurityJaqRpFetchmaterialResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpFetchmaterialResponse struct {

    // 结果信息
    
    Data   string `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpFetchmaterialResponse struct {

    // 结果信息
    Data   string `json:"data,omitempty"`

}
