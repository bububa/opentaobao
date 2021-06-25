package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象ERP核销回调 APIRequest
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调
*/
type AlibabaTclsAelophyBillVerificateCallbackRequest struct {
    model.Params

    // 回调对象
    module   *VerificateCallbackDto 

}

func NewAlibabaTclsAelophyBillVerificateCallbackRequest() *AlibabaTclsAelophyBillVerificateCallbackRequest{
    return &AlibabaTclsAelophyBillVerificateCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.bill.verificate.callback"
}

func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyBillVerificateCallbackRequest) SetModule(module *VerificateCallbackDto) error {
    r.module = module
    r.Set("module", module)
    return nil
}

func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetModule() *VerificateCallbackDto {
    return r.module
}

