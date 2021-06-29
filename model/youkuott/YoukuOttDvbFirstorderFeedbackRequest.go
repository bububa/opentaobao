package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb首次安装订单反馈 APIRequest
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

func NewYoukuOttDvbFirstorderFeedbackRequest() *YoukuOttDvbFirstorderFeedbackRequest{
    return &YoukuOttDvbFirstorderFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttDvbFirstorderFeedbackRequest) GetApiMethodName() string {
    return "youku.ott.dvb.firstorder.feedback"
}

func (r YoukuOttDvbFirstorderFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttDvbFirstorderFeedbackRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r YoukuOttDvbFirstorderFeedbackRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *YoukuOttDvbFirstorderFeedbackRequest) SetOccureTime(occureTime int64) error {
    r.occureTime = occureTime
    r.Set("occure_time", occureTime)
    return nil
}

func (r YoukuOttDvbFirstorderFeedbackRequest) GetOccureTime() int64 {
    return r.occureTime
}

func (r *YoukuOttDvbFirstorderFeedbackRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r YoukuOttDvbFirstorderFeedbackRequest) GetType() int64 {
    return r.type
}

