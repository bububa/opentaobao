package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取最外层包装码 APIRequest
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

func NewAlibabaAlihealthTracecodesellerBillRootcodeGetRequest() *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest{
    return &AlibabaAlihealthTracecodesellerBillRootcodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.rootcode.get"
}

func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetAppCode(appCode string) error {
    r.appCode = appCode
    r.Set("app_code", appCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetAppCode() string {
    return r.appCode
}

func (r *AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillRootcodeGetRequest) GetCode() string {
    return r.code
}

