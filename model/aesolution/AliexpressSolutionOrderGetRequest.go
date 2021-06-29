package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get order list APIRequest
aliexpress.solution.order.get

Get Order List from AliExpress
*/
type AliexpressSolutionOrderGetRequest struct {
    model.Params

    // param
    param0   *OrderQuery 

}

func NewAliexpressSolutionOrderGetRequest() *AliexpressSolutionOrderGetRequest{
    return &AliexpressSolutionOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.get"
}

func (r AliexpressSolutionOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionOrderGetRequest) SetParam0(param0 *OrderQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AliexpressSolutionOrderGetRequest) GetParam0() *OrderQuery {
    return r.param0
}

