package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskFinishAPIRequest
完成轻任务 API请求
taobao.qianniu.task.finish

由任务执行者调用 */
type TaobaoQianniuTaskFinishAPIRequest struct {
	model.Params
	// 任务ID
	_taskId int64
	// 任务备注
	_memo string
}

// New
