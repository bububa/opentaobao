package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardReservefailAPIRequest
预约失败 API请求
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败 */
type TmallServicecenterWorkcardReservefailAPIRequest struct {
	model.Params
	// 核销单外部id
	_identifyTaskId int64
	// 下次联系时间
	_gmtNextContact string
	// 预约失败原因码
	_failCode int64
	// 预约失败原因描述
	_failDesc string
	// 工单id
	_workcardId int64
}

// New
