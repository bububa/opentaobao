package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFinanceOrderBackflowAPIRequest
财务订单回流 API请求
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据 */
type AlibabaWdkFinanceOrderBackflowAPIRequest struct {
	model.Params
	// 财务订单回流入参
	_financeOrderDetailRequest *FinanceOrderDetailRequest
}

// NewAlibabaWdkFinanceOrderBackflowRequest 初始化AlibabaWdkFinanceOrderBackflowAPIRequest对象
func NewAlibabaWdkFinanceOrderBackflowRequest() *AlibabaWdkFinanceOrderBackflowAPIRequest {
	return &AlibabaWdkFinanceOrderBackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.finance.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FinanceOrderDetailRequest Setter
// 财务订单回流入参
func (r *AlibabaWdkFinanceOrderBackflowAPIRequest) SetFinanceOrderDetailRequest(_financeOrderDetailRequest *FinanceOrderDetailRequest) error {
	r._financeOrderDetailRequest = _financeOrderDetailRequest
	r.Set("finance_order_detail_request", _financeOrderDetailRequest)
	return nil
}

// Get FinanceOrderDetailRequest Getter
func (r AlibabaWdkFinanceOrderBackflowAPIRequest) GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
	return r._financeOrderDetailRequest
}
