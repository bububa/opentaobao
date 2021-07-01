package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerQuerypageAPIRequest
查询工人列表 API请求
tmall.servicecenter.worker.querypage

服务商查询工人列表 */
type TmallServicecenterWorkerQuerypageAPIRequest struct {
	model.Params
	// 页码
	_pageIndex int64
}

// New
