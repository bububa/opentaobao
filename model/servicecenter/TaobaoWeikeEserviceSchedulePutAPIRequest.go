package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeEserviceSchedulePutAPIRequest
提交客服排班信息 API请求
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息 */
type TaobaoWeikeEserviceSchedulePutAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 按天排班信息
	_csSchedulings []CsSchedulingOneDayDto
}

// New
