package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradetoolbindAPIRequest 批量绑定交易工具 API请求
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
type AlibabaalihousenewhometradetoolbindAPIRequest struct {
	model.Params
	// 绑定对象
	_tradeToolBindParamDto *TradeToolBindParamDto
}

// NewAlibabaalihousenewhometradetoolbindRequest 初始化AlibabaalihousenewhometradetoolbindAPIRequest对象
func NewAlibabaalihousenewhometradetoolbindRequest() *AlibabaalihousenewhometradetoolbindAPIRequest {
	return &AlibabaalihousenewhometradetoolbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhometradetoolbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.trade.tool.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhometradetoolbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhometradetoolbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeToolBindParamDto is TradeToolBindParamDto Setter
// 绑定对象
func (r *AlibabaalihousenewhometradetoolbindAPIRequest) SetTradeToolBindParamDto(_tradeToolBindParamDto *TradeToolBindParamDto) error {
	r._tradeToolBindParamDto = _tradeToolBindParamDto
	r.Set("trade_tool_bind_param_dto", _tradeToolBindParamDto)
	return nil
}

// GetTradeToolBindParamDto TradeToolBindParamDto Getter
func (r AlibabaalihousenewhometradetoolbindAPIRequest) GetTradeToolBindParamDto() *TradeToolBindParamDto {
	return r._tradeToolBindParamDto
}
