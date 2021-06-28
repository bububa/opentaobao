package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取群信息 APIResponse
taobao.openim.tribe.gettribeinfo

获取群信息
*/
type TaobaoOpenimTribeGettribeinfoAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeGettribeinfoResponse `json:"openim_tribe_gettribeinfo_response,omitempty"` 
    TaobaoOpenimTribeGettribeinfoResponse
}

/* model for simplify = false
type TaobaoOpenimTribeGettribeinfoResponse struct {

    // 群信息
    
    TribeInfo  *struct {
        TribeInfo  *TribeInfo `json:"tribe_info,omitempty"`
    } `json:"tribe_info,omitempty"`
    

}
*/

type TaobaoOpenimTribeGettribeinfoResponse struct {

    // 群信息
    TribeInfo   *TribeInfo `json:"tribe_info,omitempty"`

}
