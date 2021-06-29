package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用中心通用服务接口 API请求
yunos.tvmbos.common.operation

应用中心相关接口的代理
*/
type YunosTvmbosCommonOperationRequest struct {
    model.Params
    // 接口名
    interfaceName   string
    // 方法名
    methodName   string
    // 入参，json格式
    parameter   string
}

// 初始化YunosTvmbosCommonOperationRequest对象
func NewYunosTvmbosCommonOperationRequest() *YunosTvmbosCommonOperationRequest{
    return &YunosTvmbosCommonOperationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvmbosCommonOperationRequest) GetApiMethodName() string {
    return "yunos.tvmbos.common.operation"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvmbosCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InterfaceName Setter
// 接口名
func (r *YunosTvmbosCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

// InterfaceName Getter
func (r YunosTvmbosCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}
// MethodName Setter
// 方法名
func (r *YunosTvmbosCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

// MethodName Getter
func (r YunosTvmbosCommonOperationRequest) GetMethodName() string {
    return r.methodName
}
// Parameter Setter
// 入参，json格式
func (r *YunosTvmbosCommonOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

// Parameter Getter
func (r YunosTvmbosCommonOperationRequest) GetParameter() string {
    return r.parameter
}
