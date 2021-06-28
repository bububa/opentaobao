package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家会员（高级查询） APIResponse
taobao.crm.members.search

会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
*/
type TaobaoCrmMembersSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_members_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 根据一定条件查询的卖家会员
    
    Members   []CrmMember `json:"members,omitempty" xml:"