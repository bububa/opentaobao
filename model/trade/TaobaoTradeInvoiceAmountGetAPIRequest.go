package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeInvoiceAmountGetAPIRequest 获取订单应开票金额 API请求
// taobao.trade.invoice.amount.get
//
// 订单应开票金额计算
type TaobaoTradeInvoiceAmountGetAPIRequest struct {
	model.Params
	// 业务订单ID
	_tid int64
}

// NewTaobaoTradeInvoiceAmountGetRequest 初始化TaobaoTradeInvoiceAmountGetAPIRequest对象
func NewTaobaoTradeInvoiceAmountGetRequest() *TaobaoTradeInvoiceAmountGetAPIRequest {
	return &TaobaoTradeInvoiceAmountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.invoice.amount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 业务订单ID
func (r *TaobaoTradeInvoiceAmountGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetTid() int64 {
	return r._tid
}
