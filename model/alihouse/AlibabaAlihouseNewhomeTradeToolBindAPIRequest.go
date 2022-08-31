package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeToolBindAPIRequest 批量绑定交易工具 API请求
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
type AlibabaAlihouseNewhomeTradeToolBindAPIRequest struct {
	model.Params
	// 绑定对象
	_tradeToolBindParamDto *TradeToolBindParamDto
}

// NewAlibabaAlihouseNewhomeTradeToolBindRequest 初始化AlibabaAlihouseNewhomeTradeToolBindAPIRequest对象
func NewAlibabaAlihouseNewhomeTradeToolBindRequest() *AlibabaAlihouseNewhomeTradeToolBindAPIRequest {
	return &AlibabaAlihouseNewhomeTradeToolBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.trade.tool.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeToolBindParamDto is TradeToolBindParamDto Setter
// 绑定对象
func (r *AlibabaAlihouseNewhomeTradeToolBindAPIRequest) SetTradeToolBindParamDto(_tradeToolBindParamDto *TradeToolBindParamDto) error {
	r._tradeToolBindParamDto = _tradeToolBindParamDto
	r.Set("trade_tool_bind_param_dto", _tradeToolBindParamDto)
	return nil
}

// GetTradeToolBindParamDto TradeToolBindParamDto Getter
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetTradeToolBindParamDto() *TradeToolBindParamDto {
	return r._tradeToolBindParamDto
}
