package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Order Receipt Info APIRequest
aliexpress.solution.order.receiptinfo.get

Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
*/
type AliexpressSolutionOrderReceiptinfoGetRequest struct {
    model.Params

    // query param
    param1   *SingleOrderQuery 

}

func NewAliexpressSolutionOrderReceiptinfoGetRequest() *AliexpressSolutionOrderReceiptinfoGetRequest{
    return &AliexpressSolutionOrderReceiptinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.order.receiptinfo.get"
}

func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionOrderReceiptinfoGetRequest) SetParam1(param1 *SingleOrderQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AliexpressSolutionOrderReceiptinfoGetRequest) GetParam1() *SingleOrderQuery {
    return r.param1
}

