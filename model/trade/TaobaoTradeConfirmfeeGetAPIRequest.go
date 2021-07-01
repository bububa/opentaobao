package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeConfirmfeeGetAPIRequest
获取交易确认收货费用 API请求
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用 */
type TaobaoTradeConfirmfeeGetAPIRequest struct {
	model.Params
	// 交易主订单或子订单ID
	_tid int64
}

// NewTaobaoTradeConfirmfeeGetRequest 初始化TaobaoTradeConfirmfeeGetAPIRequest对象
func NewTaobaoTradeConfirmfeeGetRequest() *TaobaoTradeConfirmfeeGetAPIRequest {
	return &TaobaoTradeConfirmfeeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.confirmfee.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 交易主订单或子订单ID
func (r *TaobaoTradeConfirmfeeGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoTradeConfirmfeeGetAPIRequest) GetTid() int64 {
	return r._tid
}
