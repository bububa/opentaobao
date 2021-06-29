package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-创建商品等级营销明细 APIResponse
taobao.crm.grademkt.member.detail.create

创建商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailCreateAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrademktMemberDetailCreateResponse
}

type TaobaoCrmGrademktMemberDetailCreateResponse struct {
    XMLName xml.Name `xml:"crm_grademkt_member_detail_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // json格式
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
}
