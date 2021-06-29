package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
财务订单回流 API请求
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据
*/
type AlibabaWdkFinanceOrderBackflowRequest struct {
    model.Params
    // 财务订单回流入参
    _financeOrderDetailRequest   *FinanceOrderDetailRequest
}

// 初始化AlibabaWdkFinanceOrderBackflowRequest对象
func NewAlibabaWdkFinanceOrderBackflowRequest() *AlibabaWdkFinanceOrderBackflowRequest{
    return &AlibabaWdkFinanceOrderBackflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFinanceOrderBackflowRequest) GetApiMethodName() string {
    return "alibaba.wdk.finance.order.backflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFinanceOrderBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FinanceOrderDetailRequest Setter
// 财务订单回流入参
func (r *AlibabaWdkFinanceOrderBackflowRequest) SetFinanceOrderDetailRequest(_financeOrderDetailRequest *FinanceOrderDetailRequest) error {
    r._financeOrderDetailRequest = _financeOrderDetailRequest
    r.Set("finance_order_detail_request", _financeOrderDetailRequest)
    return nil
}

// FinanceOrderDetailRequest Getter
func (r AlibabaWdkFinanceOrderBackflowRequest) GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
    return r._financeOrderDetailRequest
}
