package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家的会员（基本查询） APIResponse
taobao.crm.members.get

查询卖家的会员，进行基本的查询，返回符合条件的会员列表
*/
type TaobaoCrmMembersGetAPIResponse struct {
    model.CommonResponse
    TaobaoCrmMembersGetResponse
}

type TaobaoCrmMembersGetResponse struct {
    XMLName xml.Name `xml:"crm_members_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 根据一定条件查询到卖家的会员
    
    Members   []BasicMember `json:"members,omitempty" xml:"members>basic_member,omitempty"`
    
    
    // 记录总数
    
    TotalResult   int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`

    
}
