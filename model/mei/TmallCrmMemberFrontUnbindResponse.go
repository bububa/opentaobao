package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌会员解绑 API返回值 
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
type TmallCrmMemberFrontUnbindAPIResponse struct {
    model.CommonResponse
    TmallCrmMemberFrontUnbindResponse
}

// 品牌会员解绑 成功返回结果
type TmallCrmMemberFrontUnbindResponse struct {
    XMLName xml.Name `xml:"tmall_crm_member_front_unbind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口调用是否成功
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
