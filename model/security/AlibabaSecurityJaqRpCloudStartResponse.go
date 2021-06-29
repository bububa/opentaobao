package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云开始认证 API返回值 
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpCloudStartAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpCloudStartResponse
}

// 实人认证云开始认证 成功返回结果
type AlibabaSecurityJaqRpCloudStartResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_start_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Data   *RpStartResult `json:"data,omitempty" xml:"data,omitempty"`
}
