package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增量获取卖家会员 APIResponse
taobao.crm.members.increment.get

增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
*/
type TaobaoCrmMembersIncrementGetAPIResponse struct {
    model.CommonResponse
    TaobaoCrmMembersIncrementGetResponse
}

type TaobaoCrmMembersIncrementGetResponse struct {
    XMLName xml.Name `xml:"crm_members_increment_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回当前页的会员列表
    
    Members   []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
    
    
    // 记录的总条数
    
    TotalResult   int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`

    
}
