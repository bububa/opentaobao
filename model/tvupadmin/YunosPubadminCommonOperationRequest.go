package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内部迎客松通用服务 API请求
yunos.pubadmin.common.operation

内部迎客松通用服务
*/
type YunosPubadminCommonOperationRequest struct {
    model.Params
    // 入参json串
    _parameter   string
    // 接口名
    _interfaceName   string
    // 方法名
    _methodName   string
}

// 初始化YunosPubadminCommonOperationRequest对象
func NewYunosPubadminCommonOperationRequest() *YunosPubadminCommonOperationRequest{
    return &YunosPubadminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosPubadminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.pubadmin.common.operation"
}

// IRequest interface 方法, 获取API参数
func (r YunosPubadminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameter Setter
// 入参json串
func (r *YunosPubadminCommonOperationRequest) SetParameter(_parameter string) error {
    r._parameter = _parameter
    r.Set("parameter", _parameter)
    return nil
}

// Parameter Getter
func (r YunosPubadminCommonOperationRequest) GetParameter() string {
    return r._parameter
}
// InterfaceName Setter
// 接口名
func (r *YunosPubadminCommonOperationRequest) SetInterfaceName(_interfaceName string) error {
    r._interfaceName = _interfaceName
    r.Set("interface_name", _interfaceName)
    return nil
}

// InterfaceName Getter
func (r YunosPubadminCommonOperationRequest) GetInterfaceName() string {
    return r._interfaceName
}
// MethodName Setter
// 方法名
func (r *YunosPubadminCommonOperationRequest) SetMethodName(_methodName string) error {
    r._methodName = _methodName
    r.Set("method_name", _methodName)
    return nil
}

// MethodName Getter
func (r YunosPubadminCommonOperationRequest) GetMethodName() string {
    return r._methodName
}
