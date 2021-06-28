package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群踢出成员 APIResponse
taobao.openim.tribe.expel

OPENIM群踢出成员
*/
type TaobaoOpenimTribeExpelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeExpelResponse `json:"openim_tribe_expel_response,omitempty"` 
    TaobaoOpenimTribeExpelResponse
}

/* model for simplify = false
type TaobaoOpenimTribeExpelResponse struct {

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

}
*/

type TaobaoOpenimTribeExpelResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
