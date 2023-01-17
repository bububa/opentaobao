package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest 交易工具信息上翻 API请求
// alibaba.alihouse.newhome.trade.tool.submit
//
// 交易工具信息上翻
type AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest struct {
	model.Params
	// 对象
	_tradeToolDto *TradeToolDto
}

// NewAlibabaAlihouseNewhomeTradeToolSubmitRequest 初始化AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeTradeToolSubmitRequest() *AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.trade.tool.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeToolDto is TradeToolDto Setter
// 对象
func (r *AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest) SetTradeToolDto(_tradeToolDto *TradeToolDto) error {
	r._tradeToolDto = _tradeToolDto
	r.Set("trade_tool_dto", _tradeToolDto)
	return nil
}

// GetTradeToolDto TradeToolDto Getter
func (r AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest) GetTradeToolDto() *TradeToolDto {
	return r._tradeToolDto
}
