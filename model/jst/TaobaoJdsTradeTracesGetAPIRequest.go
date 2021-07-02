package jst

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradeTracesGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trade.traces.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradeTracesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 淘宝的订单tid
func (r *TaobaoJdsTradeTracesGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoJdsTradeTracesGetAPIRequest) GetTid() int64 {
	return r._tid
}
