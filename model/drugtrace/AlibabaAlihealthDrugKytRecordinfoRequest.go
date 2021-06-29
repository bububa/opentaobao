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
    _serviceName   string
    // 类型
    _serviceType   string
    // 输入参数
    _inputParam   string
    // 其他参数
    _otherParam   string
    // 级别
    _logLevel   string
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
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceName(_serviceName string) error {
    r._serviceName = _serviceName
    r.Set("service_name", _serviceName)
    return nil
}

// ServiceName Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceName() string {
    return r._serviceName
}
// ServiceType Setter
// 类型
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetServiceType(_serviceType string) error {
    r._serviceType = _serviceType
    r.Set("service_type", _serviceType)
    return nil
}

// ServiceType Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetServiceType() string {
    return r._serviceType
}
// InputParam Setter
// 输入参数
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetInputParam(_inputParam string) error {
    r._inputParam = _inputParam
    r.Set("input_param", _inputParam)
    return nil
}

// InputParam Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetInputParam() string {
    return r._inputParam
}
// OtherParam Setter
// 其他参数
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetOtherParam(_otherParam string) error {
    r._otherParam = _otherParam
    r.Set("other_param", _otherParam)
    return nil
}

// OtherParam Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetOtherParam() string {
    return r._otherParam
}
// LogLevel Setter
// 级别
func (r *AlibabaAlihealthDrugKytRecordinfoRequest) SetLogLevel(_logLevel string) error {
    r._logLevel = _logLevel
    r.Set("log_level", _logLevel)
    return nil
}

// LogLevel Getter
func (r AlibabaAlihealthDrugKytRecordinfoRequest) GetLogLevel() string {
    return r._logLevel
}
