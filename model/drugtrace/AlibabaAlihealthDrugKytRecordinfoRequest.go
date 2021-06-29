package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通健康检查 APIRequest
alibaba.alihealth.drug.kyt.recordinfo

快易通健康检查
*/
type AlibabaAlihealthDrugKytRecordinfoRequest struct {
    model.Params

    // 服务名
    serviceName   string 

    // 类型
    serviceType   string 

    // 输入参数
    inputParam   string 

    // 其他参数
    otherParam   string 

    // 级别
    logLevel   string 

}

func NewAlibabaAlihealthDrugKytRecordinfoRequest() *AlibabaAlihealthDrugKytRecordinfoRequest{
    return &AlibabaAlihealthDrugKytRecordinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.recordinfo"
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceName(serviceName string) error {
    r.serviceName = serviceName
    r.Set("service_name", serviceName)
    return nil
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceName() string {
    return r.serviceName
}

func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceType(serviceType string) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceType() string {
    return r.serviceType
}

func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetInputParam(inputParam string) error {
    r.inputParam = inputParam
    r.Set("input_param", inputParam)
    return nil
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetInputParam() string {
    return r.inputParam
}

func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetOtherParam(otherParam string) error {
    r.otherParam = otherParam
    r.Set("other_param", otherParam)
    return nil
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetOtherParam() string {
    return r.otherParam
}

func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetLogLevel(logLevel string) error {
    r.logLevel = logLevel
    r.Set("log_level", logLevel)
    return nil
}

func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetLogLevel() string {
    return r.logLevel
}

