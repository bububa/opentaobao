package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingtradeflowsaveAPIRequest 自动售卖机交易信息回流 API请求
// alibaba.lst.vending.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabalstvendingtradeflowsaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowDTOList []VendingTradeFlowDto
}

// NewAlibabalstvendingtradeflowsaveRequest 初始化AlibabalstvendingtradeflowsaveAPIRequest对象
func NewAlibabalstvendingtradeflowsaveRequest() *AlibabalstvendingtradeflowsaveAPIRequest {
	return &AlibabalstvendingtradeflowsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingtradeflowsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.tradeflow.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingtradeflowsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingtradeflowsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeFlowDTOList is TradeFlowDTOList Setter
// 交易流水信息
func (r *AlibabalstvendingtradeflowsaveAPIRequest) SetTradeFlowDTOList(_tradeFlowDTOList []VendingTradeFlowDto) error {
	r._tradeFlowDTOList = _tradeFlowDTOList
	r.Set("trade_flow_d_t_o_list", _tradeFlowDTOList)
	return nil
}

// GetTradeFlowDTOList TradeFlowDTOList Getter
func (r AlibabalstvendingtradeflowsaveAPIRequest) GetTradeFlowDTOList() []VendingTradeFlowDto {
	return r._tradeFlowDTOList
}
