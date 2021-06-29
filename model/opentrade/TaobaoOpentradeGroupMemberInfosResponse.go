package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购批量获取用户参团信息 APIResponse
taobao.opentrade.group.member.infos

组团购场景下，获取用户参团信息
*/
type TaobaoOpentradeGroupMemberInfosAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupMemberInfosResponse
}

type TaobaoOpentradeGroupMemberInfosResponse struct {
    XMLName xml.Name `xml:"opentrade_group_member_infos_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 参团信息
    
    Result   []GroupMemberInfoResponse `json:"result,omitempty" xml:"result>group_member_info_response,omitempty"`
    
    
}
