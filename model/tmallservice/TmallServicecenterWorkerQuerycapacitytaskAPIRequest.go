package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerQuerycapacitytaskAPIRequest
查询需求容量 API请求
tmall.servicecenter.worker.querycapacitytask

查询需求容量 */
type TmallServicecenterWorkerQuerycapacitytaskAPIRequest struct {
	model.Params
	// 查询对象
	_query *CapacityTaskQueryDto
}

// New
