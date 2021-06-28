package mei

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员积分变更 APIResponse
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
type TmallCrmMemberPointChangeAPIResponse struct {
    model.CommonResponse
    // Response *TmallCrmMemberPointChangeResponse `json:"tmall_crm_member_point_change_response,omitempty"` 
    TmallCrmMemberPointChangeResponse
}

/* model for simplify = false
type TmallCrmMemberPointChangeResponse struct {

    // 调用是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty"`
    

}
*/

type TmallCrmMemberPointChangeResponse struct {

    // 调用是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

}
