package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除指定的分组或分组下的用户 APIResponse
taobao.tmc.group.delete

删除指定的分组或分组下的用户，授权消息使用
*/
type TaobaoTmcGroupDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_group_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"