package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全验证官方应用接口 APIResponse
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用
*/
type AlibabaSecurityJaqAppOfficialVerifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_app_official_verify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *OfficialAppVerifyResponse `json:"result,omitempty" xml:"