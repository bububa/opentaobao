package mei

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌会员解绑 APIResponse
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
type TmallCrmMemberFrontUnbindAPIResponse struct {
    model.CommonResponse
    Response *TmallCrmMemberFrontUnbindResponse `json:"tmall_crm_member_front_unbind_response,omitempty"`
}

type TmallCrmMemberFrontUnbindResponse struct {

    // 接口调用是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

}
