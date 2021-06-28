package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
TOP-SPI工单任务下发接口 APIResponse
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV）
*/
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse struct {
    model.CommonResponse
    CainiaoCbossWorkplatformWorkorderTaskNotifyResponse
}

type CainiaoCbossWorkplatformWorkorderTaskNotifyResponse struct {
    XMLName xml.Name `xml:"cainiao_cboss_workplatform_workorder_task_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // response
    
    Response   *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct `json:"response,omitempty" xml:"response,omitempty"`

    
}
