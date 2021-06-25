package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员关系查询 APIResponse
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。
*/
type TaobaoCrmGrademktMemberQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGrademktMemberQueryResponse `json:"taobao_crm_grademkt_member_query_response,omitempty"`
}

type TaobaoCrmGrademktMemberQueryResponse struct {

    // json格式
    Module   string `json:"module,omitempty"`

}
