package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardSigninAPIRequest
服务商确认工人签到成功 API请求
tmall.servicecenter.workcard.signin

服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担 */
type TmallServicecenterWorkcardSigninAPIRequest struct {
	model.Params
	// 核销单外部id
	_outerId string
	// 工单id
	_workcardId int64
	// 核销单id
	_fulfilTaskId int64
}

// New
