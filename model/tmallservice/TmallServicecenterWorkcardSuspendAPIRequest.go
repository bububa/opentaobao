package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardSuspendAPIRequest
工单挂起 API请求
tmall.servicecenter.workcard.suspend

工单挂起 */
type TmallServicecenterWorkcardSuspendAPIRequest struct {
	model.Params
	// 预约时间
	_reserveServiceDate string
	// 下次联系时间
	_gmtNextContact string
	// 工单id
	_workcardId int64
	// 挂起原因类型code
	_failCode int64
	// 挂起原因描述
	_failDesc string
}

// New
