package lstvending

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingTradeflowSaveAPIRequest 自动售卖机交易信息回流 API请求
// alibaba.lst.vending.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabaLstVendingTradeflowSaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowDTOList []VendingTradeFlowDto
}

// NewAlibabaLstVendingTradeflowSaveRequest 初始化AlibabaLstVendingTradeflowSaveAPIRequest对象
func NewAlibabaLstVendingTradeflowSaveRequest() *AlibabaLstVendingTradeflowSaveAPIRequest {
	return &AlibabaLstVendingTradeflowSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingTradeflowSaveAPIRequest) Reset() {
	r._tradeFlowDTOList = r._tradeFlowDTOList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.tradeflow.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeFlowDTOList is TradeFlowDTOList Setter
// 交易流水信息
func (r *AlibabaLstVendingTradeflowSaveAPIRequest) SetTradeFlowDTOList(_tradeFlowDTOList []VendingTradeFlowDto) error {
	r._tradeFlowDTOList = _tradeFlowDTOList
	r.Set("trade_flow_d_t_o_list", _tradeFlowDTOList)
	return nil
}

// GetTradeFlowDTOList TradeFlowDTOList Getter
func (r AlibabaLstVendingTradeflowSaveAPIRequest) GetTradeFlowDTOList() []VendingTradeFlowDto {
	return r._tradeFlowDTOList
}

var poolAlibabaLstVendingTradeflowSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingTradeflowSaveRequest()
	},
}

// GetAlibabaLstVendingTradeflowSaveRequest 从 sync.Pool 获取 AlibabaLstVendingTradeflowSaveAPIRequest
func GetAlibabaLstVendingTradeflowSaveAPIRequest() *AlibabaLstVendingTradeflowSaveAPIRequest {
	return poolAlibabaLstVendingTradeflowSaveAPIRequest.Get().(*AlibabaLstVendingTradeflowSaveAPIRequest)
}

// ReleaseAlibabaLstVendingTradeflowSaveAPIRequest 将 AlibabaLstVendingTradeflowSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingTradeflowSaveAPIRequest(v *AlibabaLstVendingTradeflowSaveAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingTradeflowSaveAPIRequest.Put(v)
}
