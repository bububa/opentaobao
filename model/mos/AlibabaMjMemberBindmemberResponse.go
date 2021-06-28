package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定会员 APIResponse
alibaba.mj.member.bindmember

用于绑定喵街数字化会员
*/
type AlibabaMjMemberBindmemberAPIResponse struct {
    model.CommonResponse
    AlibabaMjMemberBindmemberResponse
}

type AlibabaMjMemberBindmemberResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_member_bindmember_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
