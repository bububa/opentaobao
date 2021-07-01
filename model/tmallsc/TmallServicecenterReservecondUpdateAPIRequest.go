package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterReservecondUpdateAPIRequest
主动预约条件更新 API请求
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件 */
type TmallServicecenterReservecondUpdateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// New
