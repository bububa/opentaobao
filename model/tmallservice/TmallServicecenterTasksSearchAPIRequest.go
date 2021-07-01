package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTasksSearchAPIRequest
查询任务类工单信息 API请求
tmall.servicecenter.tasks.search

查询任务类工单信息 */
type TmallServicecenterTasksSearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// New
