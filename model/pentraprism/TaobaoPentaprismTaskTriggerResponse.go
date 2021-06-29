package pentraprism

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推进单条任务进度 API返回值 
taobao.pentaprism.task.trigger

外网用户推进单条五棱镜任务进度
*/
type TaobaoPentaprismTaskTriggerAPIResponse struct {
    model.CommonResponse
    TaobaoPentaprismTaskTriggerResponse
}

// 推进单条任务进度 成功返回结果
type TaobaoPentaprismTaskTriggerResponse struct {
    XMLName xml.Name `xml:"pentaprism_task_trigger_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // TOP接口标准出参
    Result   *TaskResult `json:"result,omitempty" xml:"result,omitempty"`
}
