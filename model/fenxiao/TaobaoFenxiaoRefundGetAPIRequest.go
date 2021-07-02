package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRefundGetAPIRequest 查询采购单退款信息 API请求
// taobao.fenxiao.refund.get
//
// 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetAPIRequest struct {
	model.Params
	// 要查询的退款子单的id
	_subOrderId int64
	// 是否查询下游买家的退款信息
	_querySellerRefund bool
}

// NewTaobaoFenxiaoRefundGetRequest 初始化TaobaoFenxiaoRefundGetAPIRequest对象
func NewTaobaoFenxiaoRefundGetRequest() *TaobaoFenxiaoRefundGetAPIRequest {
	return &TaobaoFenxiaoRefundGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRefundGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SubOrderId Setter
// 要查询的退款子单的id
func (r *TaobaoFenxiaoRefundGetAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// Get SubOrderId Getter
func (r TaobaoFenxiaoRefundGetAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// Set is QuerySellerRefund Setter
// 是否查询下游买家的退款信息
func (r *TaobaoFenxiaoRefundGetAPIRequest) SetQuerySellerRefund(_querySellerRefund bool) error {
	r._querySellerRefund = _querySellerRefund
	r.Set("query_seller_refund", _querySellerRefund)
	return nil
}

// Get QuerySellerRefund Getter
func (r TaobaoFenxiaoRefundGetAPIRequest) GetQuerySellerRefund() bool {
	return r._querySellerRefund
}
