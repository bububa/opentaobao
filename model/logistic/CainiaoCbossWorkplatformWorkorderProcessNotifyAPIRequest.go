package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest
菜鸟工单系统的工单进度下发 API请求
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI） */
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest struct {
	model.Params
	// 服务入参
	_content *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct
}

// New
