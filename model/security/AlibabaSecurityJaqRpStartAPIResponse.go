package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证开始 API返回值 
alibaba.security.jaq.rp.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpStartAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpStartAPIResponseModel
}

// 聚安全实人认证开始 成功返回结果
type AlibabaSecurityJaqRpStartAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_start_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果信息
    Data   *RpStartResult `json:"data,omitempty" xml:"data,omitempty"`
}
