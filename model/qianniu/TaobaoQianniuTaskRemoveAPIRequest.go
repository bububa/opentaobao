package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskRemoveAPIRequest
轻任务删除接口 API请求
taobao.qianniu.task.remove

轻任务删除接口。 */
type TaobaoQianniuTaskRemoveAPIRequest struct {
	model.Params
	// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
	_metadataId int64
}

// New
