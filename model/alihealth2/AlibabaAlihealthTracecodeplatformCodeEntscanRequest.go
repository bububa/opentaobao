package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品商家扫码 API请求
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回
*/
type AlibabaAlihealthTracecodeplatformCodeEntscanRequest struct {
    model.Params
    // 药监码
    _code   string
    // 不同企业有不同的标识
    _serviceFlag   string
}

// 初始化AlibabaAlihealthTracecodeplatformCodeEntscanRequest对象
func NewAlibabaAlihealthTracecodeplatformCodeEntscanRequest() *AlibabaAlihealthTracecodeplatformCodeEntscanRequest{
    return &AlibabaAlihealthTracecodeplatformCodeEntscanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeplatform.code.entscan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 药监码
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetCode() string {
    return r._code
}
// ServiceFlag Setter
// 不同企业有不同的标识
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanRequest) SetServiceFlag(_serviceFlag string) error {
    r._serviceFlag = _serviceFlag
    r.Set("service_flag", _serviceFlag)
    return nil
}

// ServiceFlag Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetServiceFlag() string {
    return r._serviceFlag
}
