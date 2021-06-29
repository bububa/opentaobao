package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb续费之后的反馈接口 APIRequest
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

func NewYoukuOttDvbRenewFeedbackRequest() *YoukuOttDvbRenewFeedbackRequest{
    return &YoukuOttDvbRenewFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttDvbRenewFeedbackRequest) GetApiMethodName() string {
    return "youku.ott.dvb.renew.feedback"
}

func (r YoukuOttDvbRenewFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttDvbRenewFeedbackRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r YoukuOttDvbRenewFeedbackRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *YoukuOttDvbRenewFeedbackRequest) SetIsSuccess(isSuccess bool) error {
    r.isSuccess = isSuccess
    r.Set("is_success", isSuccess)
    return nil
}

func (r YoukuOttDvbRenewFeedbackRequest) GetIsSuccess() bool {
    return r.isSuccess
}

func (r *YoukuOttDvbRenewFeedbackRequest) SetFailReason(failReason string) error {
    r.failReason = failReason
    r.Set("fail_reason", failReason)
    return nil
}

func (r YoukuOttDvbRenewFeedbackRequest) GetFailReason() string {
    return r.failReason
}

