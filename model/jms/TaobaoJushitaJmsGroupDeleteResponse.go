package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS分组 APIResponse
taobao.jushita.jms.group.delete

删除ONS分组
*/
type TaobaoJushitaJmsGroupDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJushitaJmsGroupDeleteResponse `json:"taobao_jushita_jms_group_delete_response,omitempty"`
}

type TaobaoJushitaJmsGroupDeleteResponse struct {

    // 操作结果
    IsSuccess   bool `json:"is_success,omitempty"`

}
