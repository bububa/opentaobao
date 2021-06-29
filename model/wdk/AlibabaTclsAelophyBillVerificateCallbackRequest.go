package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象ERP核销回调 API请求
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调
*/
type AlibabaTclsAelophyBillVerificateCallbackRequest struct {
    model.Params
    // 回调对象
    _module   *VerificateCallbackDto
}

// 初始化AlibabaTclsAelophyBillVerificateCallbackRequest对象
func NewAlibabaTclsAelophyBillVerificateCallbackRequest() *AlibabaTclsAelophyBillVerificateCallbackRequest{
    return &AlibabaTclsAelophyBillVerificateCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.bill.verificate.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Module Setter
// 回调对象
func (r *AlibabaTclsAelophyBillVerificateCallbackRequest) SetModule(_module *VerificateCallbackDto) error {
    r._module = _module
    r.Set("module", _module)
    return nil
}

// Module Getter
func (r AlibabaTclsAelophyBillVerificateCallbackRequest) GetModule() *VerificateCallbackDto {
    return r._module
}
