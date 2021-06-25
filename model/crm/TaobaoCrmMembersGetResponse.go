package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家的会员（基本查询） APIResponse
taobao.crm.members.get

查询卖家的会员，进行基本的查询，返回符合条件的会员列表
*/
type TaobaoCrmMembersGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmMembersGetResponse `json:"taobao_crm_members_get_response,omitempty"`
}

type TaobaoCrmMembersGetResponse struct {

    // 根据一定条件查询到卖家的会员
    Members   []BasicMember `json:"members,omitempty"`

    // 记录总数
    TotalResult   int64 `json:"total_result,omitempty"`

}
