package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-等级营销活动查询 APIResponse
taobao.crm.grademkt.member.detail.query

商家通过该接口查询等级营销活动
*/
type TaobaoCrmGrademktMemberDetailQueryAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrademktMemberDetailQueryResponse
}

type TaobaoCrmGrademktMemberDetailQueryResponse struct {
    XMLName xml.Name `xml:"crm_grademkt_member_detail_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // totalCount为记录总数
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
}
