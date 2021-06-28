package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新轻任务 APIResponse
taobao.qianniu.task.update

由任务执行者调用，sub_status，tag和memo至少提供一个
*/
type TaobaoQianniuTaskUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qianniu_task_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"