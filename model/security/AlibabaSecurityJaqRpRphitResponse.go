package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全-实人认证日志打点接口 APIResponse
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpRphitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_rphit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   string `json:"data,omitempty" xml:"