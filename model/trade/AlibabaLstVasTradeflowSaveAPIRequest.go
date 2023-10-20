package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVasTradeflowSaveAPIRequest 交易信息回流 API请求
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabaLstVasTradeflowSaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowModelList *TradeFlowModel
}

// NewAlibabaLstVasTradeflowSaveRequest 初始化AlibabaLstVasTradeflowSaveAPIRequest对象
func NewAlibabaLstVasTradeflowSaveRequest() *AlibabaLstVasTradeflowSaveAPIRequest {
	return &AlibabaLstVasTradeflowSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVasTradeflowSaveAPIRequest) Reset() {
	r._tradeFlowModelList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vas.tradeflow.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeFlowModelList is TradeFlowModelList Setter
// 交易流水信息
func (r *AlibabaLstVasTradeflowSaveAPIRequest) SetTradeFlowModelList(_tradeFlowModelList *TradeFlowModel) error {
	r._tradeFlowModelList = _tradeFlowModelList
	r.Set("trade_flow_model_list", _tradeFlowModelList)
	return nil
}

// GetTradeFlowModelList TradeFlowModelList Getter
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetTradeFlowModelList() *TradeFlowModel {
	return r._tradeFlowModelList
}

var poolAlibabaLstVasTradeflowSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVasTradeflowSaveRequest()
	},
}

// GetAlibabaLstVasTradeflowSaveRequest 从 sync.Pool 获取 AlibabaLstVasTradeflowSaveAPIRequest
func GetAlibabaLstVasTradeflowSaveAPIRequest() *AlibabaLstVasTradeflowSaveAPIRequest {
	return poolAlibabaLstVasTradeflowSaveAPIRequest.Get().(*AlibabaLstVasTradeflowSaveAPIRequest)
}

// ReleaseAlibabaLstVasTradeflowSaveAPIRequest 将 AlibabaLstVasTradeflowSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVasTradeflowSaveAPIRequest(v *AlibabaLstVasTradeflowSaveAPIRequest) {
	v.Reset()
	poolAlibabaLstVasTradeflowSaveAPIRequest.Put(v)
}
