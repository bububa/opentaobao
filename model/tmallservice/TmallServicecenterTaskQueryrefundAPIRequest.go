package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskQueryrefundAPIRequest
查询任务类工单是否退款 API请求
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款 */
type TmallServicecenterTaskQueryrefundAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardList []int64
}

// New
