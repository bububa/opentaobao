package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员积分变更 API返回值 
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
type TmallCrmMemberPointChangeAPIResponse struct {
    model.CommonResponse
    TmallCrmMemberPointChangeResponse
}

// 会员积分变更 成功返回结果
type TmallCrmMemberPointChangeResponse struct {
    XMLName xml.Name `xml:"tmall_crm_member_point_change_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用是否成功
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
