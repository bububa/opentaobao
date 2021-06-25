package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
TOP-SPI工单任务下发接口 APIResponse
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV）
*/
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCbossWorkplatformWorkorderTaskNotifyResponse `json:"cainiao_cboss_workplatform_workorder_task_notify_response,omitempty"`
}

type CainiaoCbossWorkplatformWorkorderTaskNotifyResponse struct {

    // response
    Response   *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct `json:"response,omitempty"`

}
