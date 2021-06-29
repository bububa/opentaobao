package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询认证结果 APIResponse
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果
*/
type AlibabaSecurityJaqRpQueryAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpQueryResponse
}

type AlibabaSecurityJaqRpQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果信息
    
    Data   *RpAuditResultBo `json:"data,omitempty" xml:"data,omitempty"`

    
}
