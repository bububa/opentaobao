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
    Response *TaobaoTmcGroupDeleteResponse `json:"taobao_tmc_group_delete_response,omitempty"`
}

type TaobaoTmcGroupDeleteResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
