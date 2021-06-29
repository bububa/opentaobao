package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get order detail info APIRequest
aliexpress.solution.order.info.get

get order detail info
*/
type AliexpressSolutionOrderInfoGetRequest struct {
    model.Params

    // param
    param1   *OrderDetailQuery 

}

func NewAliexpressSolutionOrderInfoGetRequest() *AliexpressSolutionOrderInfoGetRequest{
    return &AliexpressSolutionOrderInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionOrderInfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.info.get"
}

func (r AliexpressSolutionOrderInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionOrderInfoGetRequest) SetParam1(param1 *OrderDetailQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AliexpressSolutionOrderInfoGetRequest) GetParam1() *OrderDetailQuery {
    return r.param1
}

