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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_member_bindmember_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1
    
    Result   *SingleResult `json:"result,omitempty" xml:"