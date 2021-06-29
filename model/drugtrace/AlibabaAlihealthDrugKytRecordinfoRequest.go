package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通健康检查 API请求
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

// 初始化AlibabaAlihealthDrugKytRecordinfoRequest对象
func NewAlibabaAlihealthDrugKytRecordinfoRequest() *AlibabaAlihealthDrugKytRecordinfoRequest{
    return &AlibabaAlihealthDrugKytRecordinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.recordinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceName Setter
// 服务名
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceName(serviceName string) error {
    r.serviceName = serviceName
    r.Set("service_name", serviceName)
    return nil
}

// ServiceName Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceName() string {
    return r.serviceName
}
// ServiceType Setter
// 类型
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceType(serviceType string) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

// ServiceType Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceType() string {
    return r.serviceType
}
// InputParam Setter
// 输入参数
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetInputParam(inputParam string) error {
    r.inputParam = inputParam
    r.Set("input_param", inputParam)
    return nil
}

// InputParam Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetInputParam() string {
    return r.inputParam
}
// OtherParam Setter
// 其他参数
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetOtherParam(otherParam string) error {
    r.otherParam = otherParam
    r.Set("other_param", otherParam)
    return nil
}

// OtherParam Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetOtherParam() string {
    return r.otherParam
}
// LogLevel Setter
// 级别
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetLogLevel(logLevel string) error {
    r.logLevel = logLevel
    r.Set("log_level", logLevel)
    return nil
}

// LogLevel Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetLogLevel() string {
    return r.logLevel
}
