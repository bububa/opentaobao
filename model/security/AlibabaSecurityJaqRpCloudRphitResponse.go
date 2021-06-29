package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务日志打点 API返回值 
alibaba.security.jaq.rp.cloud.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpCloudRphitAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqRpCloudRphitResponse
}

// 实人认证云服务日志打点 成功返回结果
type AlibabaSecurityJaqRpCloudRphitResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_rphit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // result
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
}
