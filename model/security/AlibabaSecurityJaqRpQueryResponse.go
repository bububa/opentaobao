package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询认证结果 APIResponse
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果
*/
type AlibabaSecurityJaqRpQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpQueryResponse `json:"alibaba_security_jaq_rp_query_response,omitempty"` 
    AlibabaSecurityJaqRpQueryResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpQueryResponse struct {

    // 返回结果信息
    
    Data  *struct {
        RpAuditResultBo  *RpAuditResultBo `json:"rp_audit_result_bo,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpQueryResponse struct {

    // 返回结果信息
    Data   *RpAuditResultBo `json:"data,omitempty"`

}
