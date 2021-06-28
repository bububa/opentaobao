package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群主动加入 APIResponse
taobao.openim.tribe.join

OPENIM群主动加入
*/
type TaobaoOpenimTribeJoinAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeJoinResponse `json:"openim_tribe_join_response,omitempty"` 
    TaobaoOpenimTribeJoinResponse
}

/* model for simplify = false
type TaobaoOpenimTribeJoinResponse struct {

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

}
*/

type TaobaoOpenimTribeJoinResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
