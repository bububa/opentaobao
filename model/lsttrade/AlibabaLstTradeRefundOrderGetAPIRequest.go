package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeRefundOrderGetAPIRequest 零售通退款订单查询 API请求
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
type AlibabaLstTradeRefundOrderGetAPIRequest struct {
	model.Params
	// 退款单id
	_refundId string
	// 主订单id
	_mainOrderId int64
}

// NewAlibabaLstTradeRefundOrderGetRequest 初始化AlibabaLstTradeRefundOrderGetAPIRequest对象
func NewAlibabaLstTradeRefundOrderGetRequest() *AlibabaLstTradeRefundOrderGetAPIRequest {
	return &AlibabaLstTradeRefundOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeRefundOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.refund.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeRefundOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefundId is RefundId Setter
// 退款单id
func (r *AlibabaLstTradeRefundOrderGetAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaLstTradeRefundOrderGetAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabaLstTradeRefundOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaLstTradeRefundOrderGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
