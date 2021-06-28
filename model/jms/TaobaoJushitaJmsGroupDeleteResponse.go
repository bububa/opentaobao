package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS分组 APIResponse
taobao.jushita.jms.group.delete

删除ONS分组
*/
type TaobaoJushitaJmsGroupDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jushita_jms_group_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"