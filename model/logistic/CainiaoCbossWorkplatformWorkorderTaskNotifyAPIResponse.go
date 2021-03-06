package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse TOP-SPI工单任务下发接口 API返回值
// cainiao.cboss.workplatform.workorder.task.notify
//
// TOP-SPI工单任务下发接口（菜鸟--->商家ISV）
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse struct {
	model.CommonResponse
	CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponseModel
}

// CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponseModel is TOP-SPI工单任务下发接口 成功返回结果
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cboss_workplatform_workorder_task_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Response *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct `json:"response,omitempty" xml:"response,omitempty"`
}
