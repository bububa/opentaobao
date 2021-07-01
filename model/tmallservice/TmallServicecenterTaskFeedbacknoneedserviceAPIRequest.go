package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskFeedbacknoneedserviceAPIRequest
服务商反馈无需安装工单接口 API请求
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口 */
type TmallServicecenterTaskFeedbacknoneedserviceAPIRequest struct {
	model.Params
	// 入参对象
	_param *SuspendServiceDo
}

// New
