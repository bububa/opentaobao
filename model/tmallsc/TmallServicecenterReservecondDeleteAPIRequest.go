package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterReservecondDeleteAPIRequest
删除主动预约开通条件 API请求
tmall.servicecenter.reservecond.delete

删除主动预约开通条件 */
type TmallServicecenterReservecondDeleteAPIRequest struct {
	model.Params
	// 主动预约条件删除
	_reserveOpenConditionDelDto *ReserveOpenConditionDelDto
}

// New
