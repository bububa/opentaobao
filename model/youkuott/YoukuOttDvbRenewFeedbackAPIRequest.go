package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttDvbRenewFeedbackAPIRequest
dvb续费之后的反馈接口 API请求
youku.ott.dvb.renew.feedback

dvb续费之后的反馈接口 */
type YoukuOttDvbRenewFeedbackAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 是否成功
	_isSuccess bool
	// 失败原因（可无）
	_failReason string
}

// New
