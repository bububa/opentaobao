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
type AlibabaTclsAelophyBillVerificateCallbackAPIRequest struct {
    model.Params
    // 回调对象
    _module   *VerificateCallbackDTO
}

// 初始化AlibabaTclsAelophyBillVerificateCallbackAPIRequest对象
func NewAlibabaTclsAelophyBillVerificateCallbackRequest() *AlibabaTclsAelophyBillVerificateCallbackAPIRequest{
    return &AlibabaTclsAelophyBillVerificateCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.bill.verificate.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Module Setter
// 回调对象
func (r *AlibabaTclsAelophyBillVerificateCallbackAPIRequest) SetModule(_module *VerificateCallbackDTO) error {
    r._module = _module
    r.Set("module", _module)
    return nil
}

// Module Getter
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetModule() *VerificateCallbackDTO {
    return r._module
}
