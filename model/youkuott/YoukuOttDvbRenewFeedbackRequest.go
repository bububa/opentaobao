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
    orderId   int64
    // 是否成功
    isSuccess   bool
    // 失败原因（可无）
    failReason   string
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
func (r *YoukuOttDvbRenewFeedbackRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetOrderId() int64 {
    return r.orderId
}
// IsSuccess Setter
// 是否成功
func (r *YoukuOttDvbRenewFeedbackRequest) SetIsSuccess(isSuccess bool) error {
    r.isSuccess = isSuccess
    r.Set("is_success", isSuccess)
    return nil
}

// IsSuccess Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetIsSuccess() bool {
    return r.isSuccess
}
// FailReason Setter
// 失败原因（可无）
func (r *YoukuOttDvbRenewFeedbackRequest) SetFailReason(failReason string) error {
    r.failReason = failReason
    r.Set("fail_reason", failReason)
    return nil
}

// FailReason Getter
func (r YoukuOttDvbRenewFeedbackRequest) GetFailReason() string {
    return r.failReason
}
