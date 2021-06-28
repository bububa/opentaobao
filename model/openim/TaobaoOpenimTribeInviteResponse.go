package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群邀请加入 APIResponse
taobao.openim.tribe.invite

OPENIM群邀请加入接口
*/
type TaobaoOpenimTribeInviteAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeInviteResponse
}

type TaobaoOpenimTribeInviteResponse struct {
    XMLName xml.Name `xml:"openim_tribe_invite_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`

    
}
