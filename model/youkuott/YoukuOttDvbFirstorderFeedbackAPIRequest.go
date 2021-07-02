package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttDvbFirstorderFeedbackAPIRequest dvb首次安装订单反馈 API请求
// youku.ott.dvb.firstorder.feedback
//
// dvb首次安装订单反馈
type YoukuOttDvbFirstorderFeedbackAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 反馈时间（时间戳，精确到毫秒），调用接口的时间有时候并非是反馈时间，所以增加反馈时间字段作为反馈时间
	_occureTime int64
	// 反馈类型， 200：广电接单
	_type int64
}

// NewYoukuOttDvbFirstorderFeedbackRequest 初始化YoukuOttDvbFirstorderFeedbackAPIRequest对象
func NewYoukuOttDvbFirstorderFeedbackRequest() *YoukuOttDvbFirstorderFeedbackAPIRequest {
	return &YoukuOttDvbFirstorderFeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttDvbFirstorderFeedbackAPIRequest) GetApiMethodName() string {
	return "youku.ott.dvb.firstorder.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttDvbFirstorderFeedbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *YoukuOttDvbFirstorderFeedbackAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r YoukuOttDvbFirstorderFeedbackAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is OccureTime Setter
// 反馈时间（时间戳，精确到毫秒），调用接口的时间有时候并非是反馈时间，所以增加反馈时间字段作为反馈时间
func (r *YoukuOttDvbFirstorderFeedbackAPIRequest) SetOccureTime(_occureTime int64) error {
	r._occureTime = _occureTime
	r.Set("occure_time", _occureTime)
	return nil
}

// Get OccureTime Getter
func (r YoukuOttDvbFirstorderFeedbackAPIRequest) GetOccureTime() int64 {
	return r._occureTime
}

// Set is Type Setter
// 反馈类型， 200：广电接单
func (r *YoukuOttDvbFirstorderFeedbackAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r YoukuOttDvbFirstorderFeedbackAPIRequest) GetType() int64 {
	return r._type
}
