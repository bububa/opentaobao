package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterReservecondCreateAPIRequest
创建主动预约开通条件 API请求
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件 */
type TmallServicecenterReservecondCreateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// New
