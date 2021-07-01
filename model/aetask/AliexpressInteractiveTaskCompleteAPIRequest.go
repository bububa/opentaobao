package aetask

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressInteractiveTaskCompleteAPIRequest
任务完成接口 API请求
aliexpress.interactive.task.complete

用户完成任务 */
type AliexpressInteractiveTaskCompleteAPIRequest struct {
	model.Params
	// 任务实例id
	_taskInstanceId int64
	// appkey
	_projectAppKey string
}

// New
