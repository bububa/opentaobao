package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
财务订单回流 APIRequest
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据
*/
type AlibabaWdkFinanceOrderBackflowRequest struct {
    model.Params

    // 财务订单回流入参
    financeOrderDetailRequest   *FinanceOrderDetailRequest 

}

func NewAlibabaWdkFinanceOrderBackflowRequest() *AlibabaWdkFinanceOrderBackflowRequest{
    return &AlibabaWdkFinanceOrderBackflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFinanceOrderBackflowRequest) GetApiMethodName() string {
    return "alibaba.wdk.finance.order.backflow"
}

func (r AlibabaWdkFinanceOrderBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFinanceOrderBackflowRequest) SetFinanceOrderDetailRequest(financeOrderDetailRequest *FinanceOrderDetailRequest) error {
    r.financeOrderDetailRequest = financeOrderDetailRequest
    r.Set("finance_order_detail_request", financeOrderDetailRequest)
    return nil
}

func (r AlibabaWdkFinanceOrderBackflowRequest) GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
    return r.financeOrderDetailRequest
}

