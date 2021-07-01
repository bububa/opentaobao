package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingTradeflowSaveAPIRequest
自动售卖机交易信息回流 API请求
alibaba.lst.vending.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。 */
type AlibabaLstVendingTradeflowSaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowDTOList []VendingTradeFlowDto
}

// NewAlibabaLstVendingTradeflowSaveRequest 初始化AlibabaLstVendingTradeflowSaveAPIRequest对象
func NewAlibabaLstVendingTradeflowSaveRequest() *AlibabaLstVendingTradeflowSaveAPIRequest {
	return &AlibabaLstVendingTradeflowSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.tradeflow.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeFlowDTOList Setter
// 交易流水信息
func (r *AlibabaLstVendingTradeflowSaveAPIRequest) SetTradeFlowDTOList(_tradeFlowDTOList []VendingTradeFlowDto) error {
	r._tradeFlowDTOList = _tradeFlowDTOList
	r.Set("trade_flow_d_t_o_list", _tradeFlowDTOList)
	return nil
}

// Get TradeFlowDTOList Getter
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetTradeFlowDTOList() []VendingTradeFlowDto {
	return r._tradeFlowDTOList
}
