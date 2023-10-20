package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorprefundreviewAPIRequest 审核退款单 API请求
// taobao.rp.refund.review
//
// 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaorprefundreviewAPIRequest struct {
	model.Params
	// 审核人姓名
	_operator string
	// 退款阶段，可选值：售中：onsale，售后：aftersale
	_refundPhase string
	// 审核留言
	_message string
	// 退款单编号
	_refundId int64
	// 退款最后更新时间，以时间戳的方式表示
	_refundVersion int64
	// 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
	_result bool
}

// NewTaobaorprefundreviewRequest 初始化TaobaorprefundreviewAPIRequest对象
func NewTaobaorprefundreviewRequest() *TaobaorprefundreviewAPIRequest {
	return &TaobaorprefundreviewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorprefundreviewAPIRequest) GetApiMethodName() string {
	return "taobao.rp.refund.review"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorprefundreviewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorprefundreviewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperator is Operator Setter
// 审核人姓名
func (r *TaobaorprefundreviewAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaorprefundreviewAPIRequest) GetOperator() string {
	return r._operator
}

// SetRefundPhase is RefundPhase Setter
// 退款阶段，可选值：售中：onsale，售后：aftersale
func (r *TaobaorprefundreviewAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaorprefundreviewAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetMessage is Message Setter
// 审核留言
func (r *TaobaorprefundreviewAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaorprefundreviewAPIRequest) GetMessage() string {
	return r._message
}

// SetRefundId is RefundId Setter
// 退款单编号
func (r *TaobaorprefundreviewAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorprefundreviewAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款最后更新时间，以时间戳的方式表示
func (r *TaobaorprefundreviewAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaorprefundreviewAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetResult is Result Setter
// 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
func (r *TaobaorprefundreviewAPIRequest) SetResult(_result bool) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TaobaorprefundreviewAPIRequest) GetResult() bool {
	return r._result
}
