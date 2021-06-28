package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证开始 APIResponse
alibaba.security.jaq.rp.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpStartAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_start_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果信息
    
    Data   *RpStartResult `json:"data,omitempty" xml:"