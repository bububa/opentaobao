package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojecttradeorderAPIRequest 旺铺楼盘和交易商品排序 API请求
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
type AlibabaalihousenewhomeprojecttradeorderAPIRequest struct {
	model.Params
	// 参数
	_tradeOrder *ProjectTradeOrderDto
}

// NewAlibabaalihousenewhomeprojecttradeorderRequest 初始化AlibabaalihousenewhomeprojecttradeorderAPIRequest对象
func NewAlibabaalihousenewhomeprojecttradeorderRequest() *AlibabaalihousenewhomeprojecttradeorderAPIRequest {
	return &AlibabaalihousenewhomeprojecttradeorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojecttradeorderAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.trade.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojecttradeorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojecttradeorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeOrder is TradeOrder Setter
// 参数
func (r *AlibabaalihousenewhomeprojecttradeorderAPIRequest) SetTradeOrder(_tradeOrder *ProjectTradeOrderDto) error {
	r._tradeOrder = _tradeOrder
	r.Set("trade_order", _tradeOrder)
	return nil
}

// GetTradeOrder TradeOrder Getter
func (r AlibabaalihousenewhomeprojecttradeorderAPIRequest) GetTradeOrder() *ProjectTradeOrderDto {
	return r._tradeOrder
}
