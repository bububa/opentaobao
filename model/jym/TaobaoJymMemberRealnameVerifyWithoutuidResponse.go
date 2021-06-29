package jym

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户实名认证 APIResponse
taobao.jym.member.realname.verify.withoutuid

开放用户实名认证接口使用
*/
type TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse struct {
    model.CommonResponse
    TaobaoJymMemberRealnameVerifyWithoutuidResponse
}

type TaobaoJymMemberRealnameVerifyWithoutuidResponse struct {
    XMLName xml.Name `xml:"jym_member_realname_verify_withoutuid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 实名认证结果
    
    Result   *TaobaoJymMemberRealnameVerifyWithoutuidResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
