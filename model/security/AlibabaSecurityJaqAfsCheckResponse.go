package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈二次验证接口 APIResponse
alibaba.security.jaq.afs.check

反欺诈二次验证接口
*/
type AlibabaSecurityJaqAfsCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_afs_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 验证结果
    
    Data   bool `json:"data,omitempty" xml:"