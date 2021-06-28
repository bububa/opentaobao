package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证上传认证信息 APIResponse
alibaba.security.jaq.rp.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpUploadAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpUploadResponse
}

type AlibabaSecurityJaqRpUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回信息
    
    Data   *RpUploadResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
