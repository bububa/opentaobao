package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证提交认证接口 APIResponse
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpSubmitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_submit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Data   *RpSubmitResult `json:"data,omitempty" xml:"