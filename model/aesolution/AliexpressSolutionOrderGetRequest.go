package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get order list API请求
aliexpress.solution.order.get

Get Order List from AliExpress
*/
type AliexpressSolutionOrderGetRequest struct {
    model.Params
    // param
    param0   *OrderQuery
}

// 初始化AliexpressSolutionOrderGetRequest对象
func NewAliexpressSolutionOrderGetRequest() *AliexpressSolutionOrderGetRequest{
    return &AliexpressSolutionOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// param
func (r *AliexpressSolutionOrderGetRequest) SetParam0(param0 *OrderQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AliexpressSolutionOrderGetRequest) GetParam0() *OrderQuery {
    return r.param0
}
