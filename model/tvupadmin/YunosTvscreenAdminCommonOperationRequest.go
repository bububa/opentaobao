package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一体机桌面通用接口 API请求
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

// 初始化YunosTvscreenAdminCommonOperationRequest对象
func NewYunosTvscreenAdminCommonOperationRequest() *YunosTvscreenAdminCommonOperationRequest{
    return &YunosTvscreenAdminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvscreenAdminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.tvscreen.admin.common.operation"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvscreenAdminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameters Setter
// 参数数组
func (r *YunosTvscreenAdminCommonOperationRequest) SetParameters(parameters string) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

// Parameters Getter
func (r YunosTvscreenAdminCommonOperationRequest) GetParameters() string {
    return r.parameters
}
// MethodName Setter
// 方法名
func (r *YunosTvscreenAdminCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

// MethodName Getter
func (r YunosTvscreenAdminCommonOperationRequest) GetMethodName() string {
    return r.methodName
}
// InterfaceName Setter
// 接口名称
func (r *YunosTvscreenAdminCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

// InterfaceName Getter
func (r YunosTvscreenAdminCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}
