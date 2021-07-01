package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest
TOP-SPI工单任务下发接口 API请求
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV） */
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest struct {
	model.Params
	// content
	_content *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct
}

// New
