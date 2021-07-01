package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskCancelAPIRequest
取消轻任务 API请求
taobao.qianniu.task.cancel

由任务发起者调用 */
type TaobaoQianniuTaskCancelAPIRequest struct {
	model.Params
	// 任务元数据ID
	_metaId int64
	// 任务备注
	_memo string
}

// New
