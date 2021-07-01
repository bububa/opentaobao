package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskGetAPIRequest
服务工单拉取 API请求
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单 */
type TmallServicecenterTaskGetAPIRequest struct {
	model.Params
	// Taobao主交易订单ID
	_parentBizOrderId int64
}

// New
