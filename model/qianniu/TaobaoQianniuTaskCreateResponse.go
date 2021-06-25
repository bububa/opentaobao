package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建轻任务 APIResponse
taobao.qianniu.task.create

发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
*/
type TaobaoQianniuTaskCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskCreateResponse `json:"taobao_qianniu_task_create_response,omitempty"`
}

type TaobaoQianniuTaskCreateResponse struct {

    // 创建的任务元数据
    Result   *QTaskMetadata `json:"result,omitempty"`

}
