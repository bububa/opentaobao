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
    appCode   string
    // 码
    code   string
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
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetAppCode(appCode string) error {
    r.appCode = appCode
    r.Set("app_code", appCode)
    return nil
}

// AppCode Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetAppCode() string {
    return r.appCode
}
// Code Setter
// 码
func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetCode() string {
    return r.code
}
