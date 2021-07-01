package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttDvbFirstorderFeedbackAPIRequest
dvb首次安装订单反馈 API请求
youku.ott.dvb.firstorder.feedback

dvb首次安装订单反馈 */
type YoukuOttDvbFirstorderFeedbackAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 反馈时间（时间戳，精确到毫秒），调用接口的时间有时候并非是反馈时间，所以增加反馈时间字段作为反馈时间
	_occureTime int64
	// 反馈类型， 200：广电接单
	_type int64
}

// New
