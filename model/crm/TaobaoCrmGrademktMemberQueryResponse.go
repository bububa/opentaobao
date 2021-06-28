package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员关系查询 APIResponse
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。
*/
type TaobaoCrmGrademktMemberQueryAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrademktMemberQueryResponse
}

type TaobaoCrmGrademktMemberQueryResponse struct {
    XMLName xml.Name `xml:"crm_grademkt_member_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // json格式
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`

    
}
