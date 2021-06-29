package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内部迎客松通用服务 APIRequest
yunos.pubadmin.common.operation

内部迎客松通用服务
*/
type YunosPubadminCommonOperationRequest struct {
    model.Params

    // 入参json串
    parameter   string 

    // 接口名
    interfaceName   string 

    // 方法名
    methodName   string 

}

func NewYunosPubadminCommonOperationRequest() *YunosPubadminCommonOperationRequest{
    return &YunosPubadminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

func (r YunosPubadminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.pubadmin.common.operation"
}

func (r YunosPubadminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosPubadminCommonOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r YunosPubadminCommonOperationRequest) GetParameter() string {
    return r.parameter
}

func (r *YunosPubadminCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

func (r YunosPubadminCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}

func (r *YunosPubadminCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

func (r YunosPubadminCommonOperationRequest) GetMethodName() string {
    return r.methodName
}

