package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈二次验证接口 APIResponse
alibaba.security.jaq.afs.check

反欺诈二次验证接口
*/
type AlibabaSecurityJaqAfsCheckAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAfsCheckResponse
}

type AlibabaSecurityJaqAfsCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_afs_check_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 验证结果
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
}
