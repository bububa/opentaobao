package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取最外层包装码 API请求
alibaba.alihealth.tracecodeseller.bill.rootcode.get

获取最外层包装码
*/
type AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest struct {
    model.Params
    // 用户身份认证
    _appCode   string
    // 码
    _code   string
}

// 初始化AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest对象
func NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest() *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest{
    return &AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.rootcode.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppCode Setter
// 用户身份认证
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) SetAppCode(_appCode string) error {
    r._appCode = _appCode
    r.Set("app_code", _appCode)
    return nil
}

// AppCode Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetAppCode() string {
    return r._appCode
}
// Code Setter
// 码
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest) GetCode() string {
    return r._code
}
