package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易迎客松通用服务接口 API请求
yunos.trade.admin.common.operation

迎客松交易相关通用接口
*/
type YunosTradeAdminCommonOperationRequest struct {
    model.Params
    // 入参数据，json格式
    parameter   string
    // 调用方法
    methodName   string
    // 调用接口
    interfaceName   string
}

// 初始化YunosTradeAdminCommonOperationRequest对象
func NewYunosTradeAdminCommonOperationRequest() *YunosTradeAdminCommonOperationRequest{
    return &YunosTradeAdminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTradeAdminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.trade.admin.common.operation"
}

// IRequest interface 方法, 获取API参数
func (r YunosTradeAdminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameter Setter
// 入参数据，json格式
func (r *YunosTradeAdminCommonOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

// Parameter Getter
func (r YunosTradeAdminCommonOperationRequest) GetParameter() string {
    return r.parameter
}
// MethodName Setter
// 调用方法
func (r *YunosTradeAdminCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

// MethodName Getter
func (r YunosTradeAdminCommonOperationRequest) GetMethodName() string {
    return r.methodName
}
// InterfaceName Setter
// 调用接口
func (r *YunosTradeAdminCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

// InterfaceName Getter
func (r YunosTradeAdminCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}
