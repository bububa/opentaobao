package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品商家扫码 APIRequest
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回
*/
type AlibabaAlihealthTracecodeplatformCodeEntscanRequest struct {
    model.Params

    // 药监码
    code   string 

    // 不同企业有不同的标识
    serviceFlag   string 

}

func NewAlibabaAlihealthTracecodeplatformCodeEntscanRequest() *AlibabaAlihealthTracecodeplatformCodeEntscanRequest{
    return &AlibabaAlihealthTracecodeplatformCodeEntscanRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeplatform.code.entscan"
}

func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodeplatformCodeEntscanRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthTracecodeplatformCodeEntscanRequest) SetServiceFlag(serviceFlag string) error {
    r.serviceFlag = serviceFlag
    r.Set("service_flag", serviceFlag)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeEntscanRequest) GetServiceFlag() string {
    return r.serviceFlag
}

