package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsTradeTracesGetAPIRequest 获取单条订单跟踪详情 API请求
// taobao.jds.trade.traces.get
//
// 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetAPIRequest struct {
	model.Params
	// 淘宝的订单tid
	_tid int64
}

// NewTaobaoJdsTradeTracesGetRequest 初始化TaobaoJdsTradeTracesGetAPIRequest对象
func NewTaobaoJdsTradeTracesGetRequest() *TaobaoJdsTradeTracesGetAPIRequest {
	return &TaobaoJdsTradeTracesGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJdsTradeTracesGetAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradeTracesGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trade.traces.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradeTracesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJdsTradeTracesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 淘宝的订单tid
func (r *TaobaoJdsTradeTracesGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoJdsTradeTracesGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoJdsTradeTracesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJdsTradeTracesGetRequest()
	},
}

// GetTaobaoJdsTradeTracesGetRequest 从 sync.Pool 获取 TaobaoJdsTradeTracesGetAPIRequest
func GetTaobaoJdsTradeTracesGetAPIRequest() *TaobaoJdsTradeTracesGetAPIRequest {
	return poolTaobaoJdsTradeTracesGetAPIRequest.Get().(*TaobaoJdsTradeTracesGetAPIRequest)
}

// ReleaseTaobaoJdsTradeTracesGetAPIRequest 将 TaobaoJdsTradeTracesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoJdsTradeTracesGetAPIRequest(v *TaobaoJdsTradeTracesGetAPIRequest) {
	v.Reset()
	poolTaobaoJdsTradeTracesGetAPIRequest.Put(v)
}
