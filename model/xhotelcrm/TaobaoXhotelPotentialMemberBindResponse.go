package xhotelcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪酒店商家会员绑定 APIResponse
taobao.xhotel.potential.member.bind

支持互通商家发起会员绑定
*/
type TaobaoXhotelPotentialMemberBindAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelPotentialMemberBindResponse
}

type TaobaoXhotelPotentialMemberBindResponse struct {
    XMLName xml.Name `xml:"xhotel_potential_member_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 添加操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
