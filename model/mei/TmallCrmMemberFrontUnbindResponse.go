package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌会员解绑 APIResponse
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
type TmallCrmMemberFrontUnbindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_crm_member_front_unbind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口调用是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"