package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除指定的分组或分组下的用户 APIResponse
taobao.tmc.group.delete

删除指定的分组或分组下的用户，授权消息使用
*/
type TaobaoTmcGroupDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcGroupDeleteResponse `json:"tmc_group_delete_response,omitempty"` 
    TaobaoTmcGroupDeleteResponse
}

/* model for simplify = false
type TaobaoTmcGroupDeleteResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTmcGroupDeleteResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
