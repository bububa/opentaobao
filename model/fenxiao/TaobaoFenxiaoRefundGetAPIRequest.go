package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaorefundgetAPIRequest 查询采购单退款信息 API请求
// taobao.fenxiao.refund.get
//
// 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaofenxiaorefundgetAPIRequest struct {
	model.Params
	// 要查询的退款对应的分销子订单号
	_subOrderId int64
	// 退款单id（分销子订单号和退款单id，两者至少必填一个，都填的情况下，以退款单id为准）
	_refundId int64
	// 是否查询下游消费者订单对应退款信息
	_querySellerRefund bool
}

// NewTaobaofenxiaorefundgetRequest 初始化TaobaofenxiaorefundgetAPIRequest对象
func NewTaobaofenxiaorefundgetRequest() *TaobaofenxiaorefundgetAPIRequest {
	return &TaobaofenxiaorefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaorefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaorefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaorefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderId is SubOrderId Setter
// 要查询的退款对应的分销子订单号
func (r *TaobaofenxiaorefundgetAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaofenxiaorefundgetAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// SetRefundId is RefundId Setter
// 退款单id（分销子订单号和退款单id，两者至少必填一个，都填的情况下，以退款单id为准）
func (r *TaobaofenxiaorefundgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaofenxiaorefundgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetQuerySellerRefund is QuerySellerRefund Setter
// 是否查询下游消费者订单对应退款信息
func (r *TaobaofenxiaorefundgetAPIRequest) SetQuerySellerRefund(_querySellerRefund bool) error {
	r._querySellerRefund = _querySellerRefund
	r.Set("query_seller_refund", _querySellerRefund)
	return nil
}

// GetQuerySellerRefund QuerySellerRefund Getter
func (r TaobaofenxiaorefundgetAPIRequest) GetQuerySellerRefund() bool {
	return r._querySellerRefund
}
