package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradetoolsubmitAPIRequest 交易工具信息上翻 API请求
// alibaba.alihouse.newhome.trade.tool.submit
//
// 交易工具信息上翻
type AlibabaalihousenewhometradetoolsubmitAPIRequest struct {
	model.Params
	// 对象
	_tradeToolDto *TradeToolDto
}

// NewAlibabaalihousenewhometradetoolsubmitRequest 初始化AlibabaalihousenewhometradetoolsubmitAPIRequest对象
func NewAlibabaalihousenewhometradetoolsubmitRequest() *AlibabaalihousenewhometradetoolsubmitAPIRequest {
	return &AlibabaalihousenewhometradetoolsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhometradetoolsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.trade.tool.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhometradetoolsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhometradetoolsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeToolDto is TradeToolDto Setter
// 对象
func (r *AlibabaalihousenewhometradetoolsubmitAPIRequest) SetTradeToolDto(_tradeToolDto *TradeToolDto) error {
	r._tradeToolDto = _tradeToolDto
	r.Set("trade_tool_dto", _tradeToolDto)
	return nil
}

// GetTradeToolDto TradeToolDto Getter
func (r AlibabaalihousenewhometradetoolsubmitAPIRequest) GetTradeToolDto() *TradeToolDto {
	return r._tradeToolDto
}
