package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消轻任务 APIResponse
taobao.qianniu.task.cancel

由任务发起者调用
*/
type TaobaoQianniuTaskCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qianniu_task_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"