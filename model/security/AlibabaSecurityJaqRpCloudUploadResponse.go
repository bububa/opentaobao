package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云上传接口 APIResponse
alibaba.security.jaq.rp.cloud.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpCloudUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_cloud_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Data   *RpUploadResult `json:"data,omitempty" xml:"