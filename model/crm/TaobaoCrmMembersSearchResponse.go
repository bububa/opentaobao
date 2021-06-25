package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家会员（高级查询） APIResponse
taobao.crm.members.search

会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
*/
type TaobaoCrmMembersSearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmMembersSearchResponse `json:"taobao_crm_members_search_response,omitempty"`
}

type TaobaoCrmMembersSearchResponse struct {

    // 根据一定条件查询的卖家会员
    Members   []CrmMember `json:"members,omitempty"`

    // 记录的总条数
    TotalResult   int64 `json:"total_result,omitempty"`

}
