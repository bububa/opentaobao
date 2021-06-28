package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取资源文件 APIResponse
alibaba.security.jaq.resource.fetch

在前向化验证流程中提供资源文件服务
*/
type AlibabaSecurityJaqResourceFetchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_resource_fetch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取资源结果
    
    Data   *JaqResourceResult `json:"data,omitempty" xml:"