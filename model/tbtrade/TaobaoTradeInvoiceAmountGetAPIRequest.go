package tbtrade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeInvoiceAmountGetAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.invoice.amount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 业务订单ID
func (r *TaobaoTradeInvoiceAmountGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeInvoiceAmountGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoTradeInvoiceAmountGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeInvoiceAmountGetRequest()
	},
}

// GetTaobaoTradeInvoiceAmountGetRequest 从 sync.Pool 获取 TaobaoTradeInvoiceAmountGetAPIRequest
func GetTaobaoTradeInvoiceAmountGetAPIRequest() *TaobaoTradeInvoiceAmountGetAPIRequest {
	return poolTaobaoTradeInvoiceAmountGetAPIRequest.Get().(*TaobaoTradeInvoiceAmountGetAPIRequest)
}

// ReleaseTaobaoTradeInvoiceAmountGetAPIRequest 将 TaobaoTradeInvoiceAmountGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeInvoiceAmountGetAPIRequest(v *TaobaoTradeInvoiceAmountGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeInvoiceAmountGetAPIRequest.Put(v)
}
