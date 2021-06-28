package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群取消管理员 APIResponse
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员
*/
type TaobaoOpenimTribeUnsetmanagerAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeUnsetmanagerResponse `json:"openim_tribe_unsetmanager_response,omitempty"` 
    TaobaoOpenimTribeUnsetmanagerResponse
}

/* model for simplify = false
type TaobaoOpenimTribeUnsetmanagerResponse struct {

    // 群服务code
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

}
*/

type TaobaoOpenimTribeUnsetmanagerResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
