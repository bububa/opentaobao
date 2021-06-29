package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Order Receipt Info API请求
aliexpress.solution.order.receiptinfo.get

Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
*/
type AliexpressSolutionOrderReceiptinfoGetRequest struct {
    model.Params
    // query param
    _param1   *SingleOrderQuery
}

// 初始化AliexpressSolutionOrderReceiptinfoGetRequest对象
func NewAliexpressSolutionOrderReceiptinfoGetRequest() *AliexpressSolutionOrderReceiptinfoGetRequest{
    return &AliexpressSolutionOrderReceiptinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.receiptinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param1 Setter
// query param
func (r *AliexpressSolutionOrderReceiptinfoGetRequest) SetParam1(_param1 *SingleOrderQuery) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetParam1() *SingleOrderQuery {
    return r._param1
}
