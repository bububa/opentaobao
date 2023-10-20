package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeReceivetimeDelayAPIRequest 延长交易收货时间 API请求
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
type TaobaoTradeReceivetimeDelayAPIRequest struct {
	model.Params
	// 主订单号
	_tid int64
	// 延长收货的天数，可选值为：3, 5, 7, 10。
	_days int64
}

// NewTaobaoTradeReceivetimeDelayRequest 初始化TaobaoTradeReceivetimeDelayAPIRequest对象
func NewTaobaoTradeReceivetimeDelayRequest() *TaobaoTradeReceivetimeDelayAPIRequest {
	return &TaobaoTradeReceivetimeDelayAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeReceivetimeDelayAPIRequest) Reset() {
	r._tid = 0
	r._days = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetApiMethodName() string {
	return "taobao.trade.receivetime.delay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 主订单号
func (r *TaobaoTradeReceivetimeDelayAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetTid() int64 {
	return r._tid
}

// SetDays is Days Setter
// 延长收货的天数，可选值为：3, 5, 7, 10。
func (r *TaobaoTradeReceivetimeDelayAPIRequest) SetDays(_days int64) error {
	r._days = _days
	r.Set("days", _days)
	return nil
}

// GetDays Days Getter
func (r TaobaoTradeReceivetimeDelayAPIRequest) GetDays() int64 {
	return r._days
}

var poolTaobaoTradeReceivetimeDelayAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeReceivetimeDelayRequest()
	},
}

// GetTaobaoTradeReceivetimeDelayRequest 从 sync.Pool 获取 TaobaoTradeReceivetimeDelayAPIRequest
func GetTaobaoTradeReceivetimeDelayAPIRequest() *TaobaoTradeReceivetimeDelayAPIRequest {
	return poolTaobaoTradeReceivetimeDelayAPIRequest.Get().(*TaobaoTradeReceivetimeDelayAPIRequest)
}

// ReleaseTaobaoTradeReceivetimeDelayAPIRequest 将 TaobaoTradeReceivetimeDelayAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeReceivetimeDelayAPIRequest(v *TaobaoTradeReceivetimeDelayAPIRequest) {
	v.Reset()
	poolTaobaoTradeReceivetimeDelayAPIRequest.Put(v)
}
