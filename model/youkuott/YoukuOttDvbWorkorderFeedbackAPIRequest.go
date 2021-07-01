package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttDvbWorkorderFeedbackAPIRequest
dvb工单反馈 API请求
youku.ott.dvb.workorder.feedback

dvb工单处理结果反馈 */
type YoukuOttDvbWorkorderFeedbackAPIRequest struct {
	model.Params
	// 工单id
	_workorderId int64
	// 反馈内容
	_content string
	// 操作发生时间（时间戳：毫秒）
	_occureTime int64
}

// New
