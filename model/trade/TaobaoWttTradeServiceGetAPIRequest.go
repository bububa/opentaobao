package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWttTradeServiceGetAPIRequest 获取网厅号卡垂直标信息 API请求
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
type TaobaoWttTradeServiceGetAPIRequest struct {
	model.Params
	// 订单ID
	_bizOrder int64
}

// NewTaobaoWttTradeServiceGetRequest 初始化TaobaoWttTradeServiceGetAPIRequest对象
func NewTaobaoWttTradeServiceGetRequest() *TaobaoWttTradeServiceGetAPIRequest {
	return &TaobaoWttTradeServiceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWttTradeServiceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wtt.trade.service.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWttTradeServiceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrder Setter
// 订单ID
func (r *TaobaoWttTradeServiceGetAPIRequest) SetBizOrder(_bizOrder int64) error {
	r._bizOrder = _bizOrder
	r.Set("biz_order", _bizOrder)
	return nil
}

// Get BizOrder Getter
func (r TaobaoWttTradeServiceGetAPIRequest) GetBizOrder() int64 {
	return r._bizOrder
}
