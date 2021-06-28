package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群邀请加入 APIResponse
taobao.openim.tribe.invite

OPENIM群邀请加入接口
*/
type TaobaoOpenimTribeInviteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeInviteResponse `json:"openim_tribe_invite_response,omitempty"` 
    TaobaoOpenimTribeInviteResponse
}

/* model for simplify = false
type TaobaoOpenimTribeInviteResponse struct {

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

}
*/

type TaobaoOpenimTribeInviteResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
