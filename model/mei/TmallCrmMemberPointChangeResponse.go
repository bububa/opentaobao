package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员积分变更 APIResponse
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
type TmallCrmMemberPointChangeAPIResponse struct {
    model.CommonResponse
    TmallCrmMemberPointChangeResponse
}

type TmallCrmMemberPointChangeResponse struct {
    XMLName xml.Name `xml:"tmall_crm_member_point_change_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
