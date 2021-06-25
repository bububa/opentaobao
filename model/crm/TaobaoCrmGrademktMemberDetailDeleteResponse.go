package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-删除商品等级营销明细 APIResponse
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGrademktMemberDetailDeleteResponse `json:"taobao_crm_grademkt_member_detail_delete_response,omitempty"`
}

type TaobaoCrmGrademktMemberDetailDeleteResponse struct {

    // 操作是否成功
    Module   bool `json:"module,omitempty"`

}
