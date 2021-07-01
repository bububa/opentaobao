package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest
IVS事件处理反馈接口 API请求
alibaba.campus.device.openapi.feedbackeventinfo

提供给第三方ISV的的事件信息处理反馈的接口 */
type AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest struct {
	model.Params
	// 系统上下文
	_param0 *WorkBenchContext
	// 请求封装类
	_param1 *EventInfoApiDto
}

// New
