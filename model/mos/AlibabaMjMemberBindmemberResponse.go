package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
绑定会员 APIResponse
alibaba.mj.member.bindmember

用于绑定喵街数字化会员
*/
type AlibabaMjMemberBindmemberAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMemberBindmemberResponse `json:"alibaba_mj_member_bindmember_response,omitempty"`
}

type AlibabaMjMemberBindmemberResponse struct {

    // 1
    Result   *SingleResult `json:"result,omitempty"`

}
