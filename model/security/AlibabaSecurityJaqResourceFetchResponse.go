package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取资源文件 APIResponse
alibaba.security.jaq.resource.fetch

在前向化验证流程中提供资源文件服务
*/
type AlibabaSecurityJaqResourceFetchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqResourceFetchResponse `json:"alibaba_security_jaq_resource_fetch_response,omitempty"`
}

type AlibabaSecurityJaqResourceFetchResponse struct {

    // 获取资源结果
    Data   *JaqResourceResult `json:"data,omitempty"`

}
