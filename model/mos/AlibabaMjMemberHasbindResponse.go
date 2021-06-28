package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
喵街会员是否绑定 APIResponse
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员
*/
type AlibabaMjMemberHasbindAPIResponse struct {
    model.CommonResponse
    AlibabaMjMemberHasbindResponse
}

type AlibabaMjMemberHasbindResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_member_hasbind_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
