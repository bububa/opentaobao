package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增量获取卖家会员 APIResponse
taobao.crm.members.increment.get

增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
*/
type TaobaoCrmMembersIncrementGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmMembersIncrementGetResponse `json:"taobao_crm_members_increment_get_response,omitempty"`
}

type TaobaoCrmMembersIncrementGetResponse struct {

    // 返回当前页的会员列表
    Members   []BasicMember `json:"members,omitempty"`

    // 记录的总条数
    TotalResult   int64 `json:"total_result,omitempty"`

}
