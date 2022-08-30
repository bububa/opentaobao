package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest 旺铺楼盘和交易商品排序 API请求
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
type AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest struct {
	model.Params
	// 参数
	_tradeOrder *ProjectTradeOrderDto
}

// NewAlibabaAlihouseNewhomeProjectTradeOrderRequest 初始化AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectTradeOrderRequest() *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest {
	return &AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.trade.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeOrder is TradeOrder Setter
// 参数
func (r *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) SetTradeOrder(_tradeOrder *ProjectTradeOrderDto) error {
	r._tradeOrder = _tradeOrder
	r.Set("trade_order", _tradeOrder)
	return nil
}

// GetTradeOrder TradeOrder Getter
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetTradeOrder() *ProjectTradeOrderDto {
	return r._tradeOrder
}
