package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardAssignworkerAPIRequest
服务商分派工人 API请求
tmall.servicecenter.workcard.assignworker

服务商调用该接口分派工人给具体的工单 */
type TmallServicecenterWorkcardAssignworkerAPIRequest struct {
	model.Params
	// 需要指派的工人id
	_targetWorkerId int64
	// 需要指派的工人手机
	_targetWorkerMobile string
	// 需要派工人的工单id
	_workcardId int64
	// 需要指派的工人姓名
	_targetWorkerName string
	// 核销单外部id
	_outerId string
	// 不检查订单类型的原因ID由运营提供，服务商不可自由传
	_stopOrderTypeCheckReason int64
	// 核销单id
	_fulfilTaskId int64
	// 扩展信息
	_extJson string
}

// New
