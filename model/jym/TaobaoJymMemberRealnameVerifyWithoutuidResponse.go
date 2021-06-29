package jym

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户实名认证 API返回值 
taobao.jym.member.realname.verify.withoutuid

开放用户实名认证接口使用
*/
type TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse struct {
    model.CommonResponse
    TaobaoJymMemberRealnameVerifyWithoutuidResponse
}

// 用户实名认证 成功返回结果
type TaobaoJymMemberRealnameVerifyWithoutuidResponse struct {
    XMLName xml.Name `xml:"jym_member_realname_verify_withoutuid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 实名认证结果
    Result   *TaobaoJymMemberRealnameVerifyWithoutuidResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
