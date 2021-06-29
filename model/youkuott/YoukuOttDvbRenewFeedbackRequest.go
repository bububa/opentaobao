package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb续费之后的反馈接口 API请求
youku.ott.dvb.renew.feedback

dvb续费之后的反馈接口
*/
type YoukuOttDvbRenewFeedbackRequest struct {
    model.Params
    // 订单id
    _orderId   int64
    // 是否成功
    _isSuccess   bool
    // 失败原因（可无）
    _failReason   string
}

// 初始化YoukuOttDvbRenewFeedbackRequest对象
func NewYoukuOttDvbRenewFeedbackRequest() *YoukuOttDvbRenewFeedbackRequest{
    return &YoukuOttDvbRenewFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttDvbRenewFeedbackRequest) GetApiMethodName() string {
    return "youku.ott.dvb.renew.feedback"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttDvbRenewFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *YoukuOttDvbRenewFeedbackRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetOrderId() int64 {
    return r._orderId
}
// IsSuccess Setter
// 是否成功
func (r *YoukuOttDvbRenewFeedbackRequest) SetIsSuccess(_isSuccess bool) error {
    r._isSuccess = _isSuccess
    r.Set("is_success", _isSuccess)
    return nil
}

// IsSuccess Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetIsSuccess() bool {
    return r._isSuccess
}
// FailReason Setter
// 失败原因（可无）
func (r *YoukuOttDvbRenewFeedbackRequest) SetFailReason(_failReason string) error {
    r._failReason = _failReason
    r.Set("fail_reason", _failReason)
    return nil
}

// FailReason Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetFailReason() string {
    return r._failReason
}
