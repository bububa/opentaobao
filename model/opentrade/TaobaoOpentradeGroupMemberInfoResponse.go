package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取用户参团信息 APIResponse
taobao.opentrade.group.member.info

组团购场景下，获取用户参团信息
*/
type TaobaoOpentradeGroupMemberInfoAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupMemberInfoResponse
}

type TaobaoOpentradeGroupMemberInfoResponse struct {
    XMLName xml.Name `xml:"opentrade_group_member_info_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 参团信息
    
    Result   *RopGroupMemberInfo `json:"result,omitempty" xml:"result,omitempty"`

    
}
