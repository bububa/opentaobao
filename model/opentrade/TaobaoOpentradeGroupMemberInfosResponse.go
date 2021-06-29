package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购批量获取用户参团信息 API返回值 
taobao.opentrade.group.member.infos

组团购场景下，获取用户参团信息
*/
type TaobaoOpentradeGroupMemberInfosAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupMemberInfosResponse
}

// 组团购批量获取用户参团信息 成功返回结果
type TaobaoOpentradeGroupMemberInfosResponse struct {
    XMLName xml.Name `xml:"opentrade_group_member_infos_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 参团信息
    Result   []GroupMemberInfoResponse `json:"result,omitempty" xml:"result>group_member_info_response,omitempty"`
}
