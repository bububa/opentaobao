package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttDvbRenewFeedbackAPIRequest dvb续费之后的反馈接口 API请求
// youku.ott.dvb.renew.feedback
//
// dvb续费之后的反馈接口
type YoukuOttDvbRenewFeedbackAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 是否成功
	_isSuccess bool
	// 失败原因（可无）
	_failReason string
}

// NewYoukuOttDvbRenewFeedbackRequest 初始化YoukuOttDvbRenewFeedbackAPIRequest对象
func NewYoukuOttDvbRenewFeedbackRequest() *YoukuOttDvbRenewFeedbackAPIRequest {
	return &YoukuOttDvbRenewFeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttDvbRenewFeedbackAPIRequest) GetApiMethodName() string {
	return "youku.ott.dvb.renew.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttDvbRenewFeedbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *YoukuOttDvbRenewFeedbackAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r YoukuOttDvbRenewFeedbackAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetIsSuccess is IsSuccess Setter
// 是否成功
func (r *YoukuOttDvbRenewFeedbackAPIRequest) SetIsSuccess(_isSuccess bool) error {
	r._isSuccess = _isSuccess
	r.Set("is_success", _isSuccess)
	return nil
}

// GetIsSuccess IsSuccess Getter
func (r YoukuOttDvbRenewFeedbackAPIRequest) GetIsSuccess() bool {
	return r._isSuccess
}

// SetFailReason is FailReason Setter
// 失败原因（可无）
func (r *YoukuOttDvbRenewFeedbackAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// GetFailReason FailReason Getter
func (r YoukuOttDvbRenewFeedbackAPIRequest) GetFailReason() string {
	return r._failReason
}
