package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-删除商品等级营销明细 API返回值 
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrademktMemberDetailDeleteResponse
}

// 会员等级营销-删除商品等级营销明细 成功返回结果
type TaobaoCrmGrademktMemberDetailDeleteResponse struct {
    XMLName xml.Name `xml:"crm_grademkt_member_detail_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作是否成功
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
}
