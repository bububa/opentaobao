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
type AlibabaAlihealthTracecodesellerBillRootcodeGetRequest struct {
    model.Params
    // 用户身份认证
    _appCode   string
    // 码
    _code   string
}

// 初始化AlibabaAlihealthTracecodesellerBillRootcodeGetRequest对象
func NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest() *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest{
    return &AlibabaAlihealthTracecodesellerBillRootcodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.rootcode.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppCode Setter
// 用户身份认证
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetAppCode(_appCode string) error {
    r._appCode = _appCode
    r.Set("app_code", _appCode)
    return nil
}

// AppCode Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetAppCode() string {
    return r._appCode
}
// Code Setter
// 码
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetCode() string {
    return r._code
}
