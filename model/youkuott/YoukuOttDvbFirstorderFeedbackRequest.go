package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb首次安装订单反馈 API请求
youku.ott.dvb.firstorder.feedback

dvb首次安装订单反馈
*/
type YoukuOttDvbFirstorderFeedbackRequest struct {
    model.Params
    // 订单id
    orderId   int64
    // 反馈时间（时间戳，精确到毫秒），调用接口的时间有时候并非是反馈时间，所以增加反馈时间字段作为反馈时间
    occureTime   int64
    // 反馈类型， 200：广电接单
    type   int64
}

// 初始化YoukuOttDvbFirstorderFeedbackRequest对象
func NewYoukuOttDvbFirstorderFeedbackRequest() *YoukuOttDvbFirstorderFeedbackRequest{
    return &YoukuOttDvbFirstorderFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttDvbFirstorderFeedbackRequest) GetApiMethodName() string {
    return "youku.ott.dvb.firstorder.feedback"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttDvbFirstorderFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *YoukuOttDvbFirstorderFeedbackRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r YoukuOttDvbFirstorderFeedbackRequest) GetOrderId() int64 {
    return r.orderId
}
// OccureTime Setter
// 反馈时间（时间戳，精确到毫秒），调用接口的时间有时候并非是反馈时间，所以增加反馈时间字段作为反馈时间
func (r *YoukuOttDvbFirstorderFeedbackRequest) SetOccureTime(occureTime int64) error {
    r.occureTime = occureTime
    r.Set("occure_time", occureTime)
    return nil
}

// OccureTime Getter
func (r YoukuOttDvbFirstorderFeedbackRequest) GetOccureTime() int64 {
    return r.occureTime
}
// Type Setter
// 反馈类型， 200：广电接单
func (r *YoukuOttDvbFirstorderFeedbackRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r YoukuOttDvbFirstorderFeedbackRequest) GetType() int64 {
    return r.type
}
