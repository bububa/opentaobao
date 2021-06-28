package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
将一个分组添加到另外一个分组 APIResponse
taobao.crm.group.append

将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。
*/
type TaobaoCrmGroupAppendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_group_append_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步任务请求成功，添加任务是否完成通过taobao.crm.grouptask.check检测
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"