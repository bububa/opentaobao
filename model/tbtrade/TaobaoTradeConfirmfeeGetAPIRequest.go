package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeConfirmfeeGetAPIRequest 获取交易确认收货费用 API请求
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaoTradeConfirmfeeGetAPIRequest struct {
	model.Params
	// 交易主订单或子订单ID
	_tid int64
}

// NewTaobaoTradeConfirmfeeGetRequest 初始化TaobaoTradeConfirmfeeGetAPIRequest对象
func NewTaobaoTradeConfirmfeeGetRequest() *TaobaoTradeConfirmfeeGetAPIRequest {
	return &TaobaoTradeConfirmfeeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeConfirmfeeGetAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.confirmfee.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易主订单或子订单ID
func (r *TaobaoTradeConfirmfeeGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoTradeConfirmfeeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeConfirmfeeGetRequest()
	},
}

// GetTaobaoTradeConfirmfeeGetRequest 从 sync.Pool 获取 TaobaoTradeConfirmfeeGetAPIRequest
func GetTaobaoTradeConfirmfeeGetAPIRequest() *TaobaoTradeConfirmfeeGetAPIRequest {
	return poolTaobaoTradeConfirmfeeGetAPIRequest.Get().(*TaobaoTradeConfirmfeeGetAPIRequest)
}

// ReleaseTaobaoTradeConfirmfeeGetAPIRequest 将 TaobaoTradeConfirmfeeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeConfirmfeeGetAPIRequest(v *TaobaoTradeConfirmfeeGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeConfirmfeeGetAPIRequest.Put(v)
}
