package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-创建商品等级营销明细 APIResponse
taobao.crm.grademkt.member.detail.create

创建商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGrademktMemberDetailCreateResponse `json:"taobao_crm_grademkt_member_detail_create_response,omitempty"`
}

type TaobaoCrmGrademktMemberDetailCreateResponse struct {

    // json格式
    Module   bool `json:"module,omitempty"`

}
