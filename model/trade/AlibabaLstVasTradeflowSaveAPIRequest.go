package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvastradeflowsaveAPIRequest 交易信息回流 API请求
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabalstvastradeflowsaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowModelList *TradeFlowModel
}

// NewAlibabalstvastradeflowsaveRequest 初始化AlibabalstvastradeflowsaveAPIRequest对象
func NewAlibabalstvastradeflowsaveRequest() *AlibabalstvastradeflowsaveAPIRequest {
	return &AlibabalstvastradeflowsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvastradeflowsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vas.tradeflow.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvastradeflowsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvastradeflowsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeFlowModelList is TradeFlowModelList Setter
// 交易流水信息
func (r *AlibabalstvastradeflowsaveAPIRequest) SetTradeFlowModelList(_tradeFlowModelList *TradeFlowModel) error {
	r._tradeFlowModelList = _tradeFlowModelList
	r.Set("trade_flow_model_list", _tradeFlowModelList)
	return nil
}

// GetTradeFlowModelList TradeFlowModelList Getter
func (r AlibabalstvastradeflowsaveAPIRequest) GetTradeFlowModelList() *TradeFlowModel {
	return r._tradeFlowModelList
}
