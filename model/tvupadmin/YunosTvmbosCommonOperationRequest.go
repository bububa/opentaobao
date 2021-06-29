package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用中心通用服务接口 APIRequest
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

func NewYunosTvmbosCommonOperationRequest() *YunosTvmbosCommonOperationRequest{
    return &YunosTvmbosCommonOperationRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvmbosCommonOperationRequest) GetApiMethodName() string {
    return "yunos.tvmbos.common.operation"
}

func (r YunosTvmbosCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvmbosCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

func (r YunosTvmbosCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}

func (r *YunosTvmbosCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

func (r YunosTvmbosCommonOperationRequest) GetMethodName() string {
    return r.methodName
}

func (r *YunosTvmbosCommonOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r YunosTvmbosCommonOperationRequest) GetParameter() string {
    return r.parameter
}

