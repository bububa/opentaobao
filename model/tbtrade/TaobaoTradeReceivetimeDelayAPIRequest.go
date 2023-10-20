package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradereceivetimedelayAPIRequest 延长交易收货时间 API请求
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
type TaobaotradereceivetimedelayAPIRequest struct {
	model.Params
	// 主订单号
	_tid int64
	// 延长收货的天数，可选值为：3, 5, 7, 10。
	_days int64
}

// NewTaobaotradereceivetimedelayRequest 初始化TaobaotradereceivetimedelayAPIRequest对象
func NewTaobaotradereceivetimedelayRequest() *TaobaotradereceivetimedelayAPIRequest {
	return &TaobaotradereceivetimedelayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradereceivetimedelayAPIRequest) GetApiMethodName() string {
	return "taobao.trade.receivetime.delay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradereceivetimedelayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradereceivetimedelayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 主订单号
func (r *TaobaotradereceivetimedelayAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradereceivetimedelayAPIRequest) GetTid() int64 {
	return r._tid
}

// SetDays is Days Setter
// 延长收货的天数，可选值为：3, 5, 7, 10。
func (r *TaobaotradereceivetimedelayAPIRequest) SetDays(_days int64) error {
	r._days = _days
	r.Set("days", _days)
	return nil
}

// GetDays Days Getter
func (r TaobaotradereceivetimedelayAPIRequest) GetDays() int64 {
	return r._days
}
