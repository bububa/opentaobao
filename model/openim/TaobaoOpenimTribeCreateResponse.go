package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建群 APIResponse
taobao.openim.tribe.create

创建一个openim的群
*/
type TaobaoOpenimTribeCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeCreateResponse `json:"openim_tribe_create_response,omitempty"` 
    TaobaoOpenimTribeCreateResponse
}

/* model for simplify = false
type TaobaoOpenimTribeCreateResponse struct {

    // 创建群的信息
    
    TribeInfo  *struct {
        TribeInfo  *TribeInfo `json:"tribe_info,omitempty"`
    } `json:"tribe_info,omitempty"`
    

}
*/

type TaobaoOpenimTribeCreateResponse struct {

    // 创建群的信息
    TribeInfo   *TribeInfo `json:"tribe_info,omitempty"`

}
