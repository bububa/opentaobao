package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-会员吸纳 APIResponse
taobao.crm.grademkt.member.add

商家通过该接口吸纳线上店铺会员。
*/
type TaobaoCrmGrademktMemberAddAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrademktMemberAddResponse
}

type TaobaoCrmGrademktMemberAddResponse struct {
    XMLName xml.Name `xml:"crm_grademkt_member_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回操作是否成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
