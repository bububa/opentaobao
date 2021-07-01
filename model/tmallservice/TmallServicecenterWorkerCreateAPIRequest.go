package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerCreateAPIRequest
服务商工人信息创建 API请求
tmall.servicecenter.worker.create

服务商工人信息创建 */
type TmallServicecenterWorkerCreateAPIRequest struct {
	model.Params
	// 11
	_workerDto *WorkerDto
}

// New
