package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
验证姓名和证件号 APIResponse
alibaba.security.jaq.rp.cloud.realname.check

验证姓名和证件号
*/
type AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_rp_cloud_realname_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Data   *RealNameResult `json:"data,omitempty" xml:"