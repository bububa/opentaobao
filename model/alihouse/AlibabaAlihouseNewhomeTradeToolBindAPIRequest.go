package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeTradeToolBindAPIRequest) Reset() {
	r._tradeToolBindParamDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.trade.tool.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeTradeToolBindAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeTradeToolBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeTradeToolBindRequest()
	},
}

// GetAlibabaAlihouseNewhomeTradeToolBindRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeTradeToolBindAPIRequest
func GetAlibabaAlihouseNewhomeTradeToolBindAPIRequest() *AlibabaAlihouseNewhomeTradeToolBindAPIRequest {
	return poolAlibabaAlihouseNewhomeTradeToolBindAPIRequest.Get().(*AlibabaAlihouseNewhomeTradeToolBindAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeTradeToolBindAPIRequest 将 AlibabaAlihouseNewhomeTradeToolBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTradeToolBindAPIRequest(v *AlibabaAlihouseNewhomeTradeToolBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTradeToolBindAPIRequest.Put(v)
}
