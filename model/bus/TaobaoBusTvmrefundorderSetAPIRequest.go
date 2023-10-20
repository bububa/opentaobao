package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmrefundordersetAPIRequest 线下自助机逆向退款接口 API请求
// taobao.bus.tvmrefundorder.set
//
// 汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
type TaobaobustvmrefundordersetAPIRequest struct {
	model.Params
	// 分账退款明细
	_refundAccountInDetails []RefundAccountInDetail
	// 保险退款详情
	_insuranceRefundDetails []InsuranceRefundDetail
	// 飞猪订单号
	_alitripOrderId string
	// 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
	_refundBatchNo string
	// 退款原因
	_refundReason string
	// 退款金额（单位分） 票金额
	_refundAmount int64
}

// NewTaobaobustvmrefundordersetRequest 初始化TaobaobustvmrefundordersetAPIRequest对象
func NewTaobaobustvmrefundordersetRequest() *TaobaobustvmrefundordersetAPIRequest {
	return &TaobaobustvmrefundordersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmrefundordersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmrefundorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmrefundordersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmrefundordersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundAccountInDetails is RefundAccountInDetails Setter
// 分账退款明细
func (r *TaobaobustvmrefundordersetAPIRequest) SetRefundAccountInDetails(_refundAccountInDetails []RefundAccountInDetail) error {
	r._refundAccountInDetails = _refundAccountInDetails
	r.Set("refund_account_in_details", _refundAccountInDetails)
	return nil
}

// GetRefundAccountInDetails RefundAccountInDetails Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetRefundAccountInDetails() []RefundAccountInDetail {
	return r._refundAccountInDetails
}

// SetInsuranceRefundDetails is InsuranceRefundDetails Setter
// 保险退款详情
func (r *TaobaobustvmrefundordersetAPIRequest) SetInsuranceRefundDetails(_insuranceRefundDetails []InsuranceRefundDetail) error {
	r._insuranceRefundDetails = _insuranceRefundDetails
	r.Set("insurance_refund_details", _insuranceRefundDetails)
	return nil
}

// GetInsuranceRefundDetails InsuranceRefundDetails Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetInsuranceRefundDetails() []InsuranceRefundDetail {
	return r._insuranceRefundDetails
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobustvmrefundordersetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

// SetRefundBatchNo is RefundBatchNo Setter
// 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
func (r *TaobaobustvmrefundordersetAPIRequest) SetRefundBatchNo(_refundBatchNo string) error {
	r._refundBatchNo = _refundBatchNo
	r.Set("refund_batch_no", _refundBatchNo)
	return nil
}

// GetRefundBatchNo RefundBatchNo Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetRefundBatchNo() string {
	return r._refundBatchNo
}

// SetRefundReason is RefundReason Setter
// 退款原因
func (r *TaobaobustvmrefundordersetAPIRequest) SetRefundReason(_refundReason string) error {
	r._refundReason = _refundReason
	r.Set("refund_reason", _refundReason)
	return nil
}

// GetRefundReason RefundReason Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetRefundReason() string {
	return r._refundReason
}

// SetRefundAmount is RefundAmount Setter
// 退款金额（单位分） 票金额
func (r *TaobaobustvmrefundordersetAPIRequest) SetRefundAmount(_refundAmount int64) error {
	r._refundAmount = _refundAmount
	r.Set("refund_amount", _refundAmount)
	return nil
}

// GetRefundAmount RefundAmount Getter
func (r TaobaobustvmrefundordersetAPIRequest) GetRefundAmount() int64 {
	return r._refundAmount
}
