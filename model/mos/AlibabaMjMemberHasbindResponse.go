package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
喵街会员是否绑定 APIResponse
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员
*/
type AlibabaMjMemberHasbindAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMemberHasbindResponse `json:"alibaba_mj_member_hasbind_response,omitempty"`
}

type AlibabaMjMemberHasbindResponse struct {

    // 结果
    Result   *SingleResult `json:"result,omitempty"`

}
