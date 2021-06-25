package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-等级营销活动查询 APIResponse
taobao.crm.grademkt.member.detail.query

商家通过该接口查询等级营销活动
*/
type TaobaoCrmGrademktMemberDetailQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGrademktMemberDetailQueryResponse `json:"taobao_crm_grademkt_member_detail_query_response,omitempty"`
}

type TaobaoCrmGrademktMemberDetailQueryResponse struct {

    // totalCount为记录总数
    Model   string `json:"model,omitempty"`

}
