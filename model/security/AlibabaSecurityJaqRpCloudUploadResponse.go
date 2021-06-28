package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云上传接口 APIResponse
alibaba.security.jaq.rp.cloud.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpCloudUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqRpCloudUploadResponse `json:"alibaba_security_jaq_rp_cloud_upload_response,omitempty"` 
    AlibabaSecurityJaqRpCloudUploadResponse
}

/* model for simplify = false
type AlibabaSecurityJaqRpCloudUploadResponse struct {

    // result
    
    Data  *struct {
        RpUploadResult  *RpUploadResult `json:"rp_upload_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqRpCloudUploadResponse struct {

    // result
    Data   *RpUploadResult `json:"data,omitempty"`

}
