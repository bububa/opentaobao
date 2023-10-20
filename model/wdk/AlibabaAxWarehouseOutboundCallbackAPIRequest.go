package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaxwarehouseoutboundcallbackAPIRequest 翱象出仓回传 API请求
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
type AlibabaaxwarehouseoutboundcallbackAPIRequest struct {
	model.Params
	// 出库回传
	_tradeOutBoundCallBackRequest *TopTradeOutBoundCallBackRequest
}

// NewAlibabaaxwarehouseoutboundcallbackRequest 初始化AlibabaaxwarehouseoutboundcallbackAPIRequest对象
func NewAlibabaaxwarehouseoutboundcallbackRequest() *AlibabaaxwarehouseoutboundcallbackAPIRequest {
	return &AlibabaaxwarehouseoutboundcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaxwarehouseoutboundcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ax.warehouse.outbound.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaxwarehouseoutboundcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaxwarehouseoutboundcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeOutBoundCallBackRequest is TradeOutBoundCallBackRequest Setter
// 出库回传
func (r *AlibabaaxwarehouseoutboundcallbackAPIRequest) SetTradeOutBoundCallBackRequest(_tradeOutBoundCallBackRequest *TopTradeOutBoundCallBackRequest) error {
	r._tradeOutBoundCallBackRequest = _tradeOutBoundCallBackRequest
	r.Set("trade_out_bound_call_back_request", _tradeOutBoundCallBackRequest)
	return nil
}

// GetTradeOutBoundCallBackRequest TradeOutBoundCallBackRequest Getter
func (r AlibabaaxwarehouseoutboundcallbackAPIRequest) GetTradeOutBoundCallBackRequest() *TopTradeOutBoundCallBackRequest {
	return r._tradeOutBoundCallBackRequest
}
