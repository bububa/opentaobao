package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFinanceOrderBackflowAPIRequest 财务订单回流 API请求
// alibaba.wdk.finance.order.backflow
//
// 星巴克拉取财务订单回流数据
type AlibabaWdkFinanceOrderBackflowAPIRequest struct {
	model.Params
	// 财务订单回流入参
	_financeOrderDetailRequest *FinanceOrderDetailRequest
}

// NewAlibabaWdkFinanceOrderBackflowRequest 初始化AlibabaWdkFinanceOrderBackflowAPIRequest对象
func NewAlibabaWdkFinanceOrderBackflowRequest() *AlibabaWdkFinanceOrderBackflowAPIRequest {
	return &AlibabaWdkFinanceOrderBackflowAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFinanceOrderBackflowAPIRequest) Reset() {
	r._financeOrderDetailRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.finance.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFinanceOrderDetailRequest is FinanceOrderDetailRequest Setter
// 财务订单回流入参
func (r *AlibabaWdkFinanceOrderBackflowAPIRequest) SetFinanceOrderDetailRequest(_financeOrderDetailRequest *FinanceOrderDetailRequest) error {
	r._financeOrderDetailRequest = _financeOrderDetailRequest
	r.Set("finance_order_detail_request", _financeOrderDetailRequest)
	return nil
}

// GetFinanceOrderDetailRequest FinanceOrderDetailRequest Getter
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
	return r._financeOrderDetailRequest
}

var poolAlibabaWdkFinanceOrderBackflowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFinanceOrderBackflowRequest()
	},
}

// GetAlibabaWdkFinanceOrderBackflowRequest 从 sync.Pool 获取 AlibabaWdkFinanceOrderBackflowAPIRequest
func GetAlibabaWdkFinanceOrderBackflowAPIRequest() *AlibabaWdkFinanceOrderBackflowAPIRequest {
	return poolAlibabaWdkFinanceOrderBackflowAPIRequest.Get().(*AlibabaWdkFinanceOrderBackflowAPIRequest)
}

// ReleaseAlibabaWdkFinanceOrderBackflowAPIRequest 将 AlibabaWdkFinanceOrderBackflowAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFinanceOrderBackflowAPIRequest(v *AlibabaWdkFinanceOrderBackflowAPIRequest) {
	v.Reset()
	poolAlibabaWdkFinanceOrderBackflowAPIRequest.Put(v)
}
