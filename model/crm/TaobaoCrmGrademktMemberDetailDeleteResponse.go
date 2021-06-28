package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-删除商品等级营销明细 APIResponse
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_grademkt_member_detail_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功
    
    Module   bool `json:"module,omitempty" xml:"