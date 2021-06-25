package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员吸纳 APIResponse
taobao.crm.grademkt.member.add

商家通过该接口吸纳线上店铺会员。
*/
type TaobaoCrmGrademktMemberAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGrademktMemberAddResponse `json:"taobao_crm_grademkt_member_add_response,omitempty"`
}

type TaobaoCrmGrademktMemberAddResponse struct {

    // 返回操作是否成功
    Model   bool `json:"model,omitempty"`

}
