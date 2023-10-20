package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdstradetracesgetAPIRequest 获取单条订单跟踪详情 API请求
// taobao.jds.trade.traces.get
//
// 获取聚石塔数据共享的交易全链路信息
type TaobaojdstradetracesgetAPIRequest struct {
	model.Params
	// 淘宝的订单tid
	_tid int64
}

// NewTaobaojdstradetracesgetRequest 初始化TaobaojdstradetracesgetAPIRequest对象
func NewTaobaojdstradetracesgetRequest() *TaobaojdstradetracesgetAPIRequest {
	return &TaobaojdstradetracesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojdstradetracesgetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.trade.traces.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojdstradetracesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojdstradetracesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 淘宝的订单tid
func (r *TaobaojdstradetracesgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaojdstradetracesgetAPIRequest) GetTid() int64 {
	return r._tid
}
