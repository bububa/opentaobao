package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取结果接口 APIResponse
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
type AlibabaSecurityJaqRpFetchmaterialAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_fetchmaterial_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Data   string `json:"data,omitempty" xml:"