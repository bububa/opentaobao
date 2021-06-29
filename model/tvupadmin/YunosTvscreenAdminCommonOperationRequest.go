package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一体机桌面通用接口 APIRequest
yunos.tvscreen.admin.common.operation

一体机桌面通用接口
*/
type YunosTvscreenAdminCommonOperationRequest struct {
    model.Params

    // 参数数组
    parameters   string 

    // 方法名
    methodName   string 

    // 接口名称
    interfaceName   string 

}

func NewYunosTvscreenAdminCommonOperationRequest() *YunosTvscreenAdminCommonOperationRequest{
    return &YunosTvscreenAdminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvscreenAdminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.tvscreen.admin.common.operation"
}

func (r YunosTvscreenAdminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvscreenAdminCommonOperationRequest) SetParameters(parameters string) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

func (r YunosTvscreenAdminCommonOperationRequest) GetParameters() string {
    return r.parameters
}

func (r *YunosTvscreenAdminCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

func (r YunosTvscreenAdminCommonOperationRequest) GetMethodName() string {
    return r.methodName
}

func (r *YunosTvscreenAdminCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

func (r YunosTvscreenAdminCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}

