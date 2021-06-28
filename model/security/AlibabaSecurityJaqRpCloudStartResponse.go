package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云开始认证 APIResponse
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpCloudStartAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_cloud_start_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Data   *RpStartResult `json:"data,omitempty" xml:"