package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询状态接口 APIResponse
alibaba.security.jaq.rp.status

聚安全实人认证查询状态接口
*/
type AlibabaSecurityJaqRpStatusAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpStatusResponse
}

type AlibabaSecurityJaqRpStatusResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_status_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 状态信息
    
    Data   *RpStatusResultBo `json:"data,omitempty" xml:"data,omitempty"`

    
}
