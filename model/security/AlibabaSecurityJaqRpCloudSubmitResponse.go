package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务提交接口 APIResponse
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpCloudSubmitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_cloud_submit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Data   *RpSubmitResult `json:"data,omitempty" xml:"