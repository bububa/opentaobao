package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradeconfirmfeegetAPIRequest 获取交易确认收货费用 API请求
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
type TaobaotradeconfirmfeegetAPIRequest struct {
	model.Params
	// 交易主订单或子订单ID
	_tid int64
}

// NewTaobaotradeconfirmfeegetRequest 初始化TaobaotradeconfirmfeegetAPIRequest对象
func NewTaobaotradeconfirmfeegetRequest() *TaobaotradeconfirmfeegetAPIRequest {
	return &TaobaotradeconfirmfeegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradeconfirmfeegetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.confirmfee.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradeconfirmfeegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradeconfirmfeegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易主订单或子订单ID
func (r *TaobaotradeconfirmfeegetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradeconfirmfeegetAPIRequest) GetTid() int64 {
	return r._tid
}
