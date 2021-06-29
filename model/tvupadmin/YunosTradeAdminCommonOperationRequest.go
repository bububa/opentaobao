package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易迎客松通用服务接口 APIRequest
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

func NewYunosTradeAdminCommonOperationRequest() *YunosTradeAdminCommonOperationRequest{
    return &YunosTradeAdminCommonOperationRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTradeAdminCommonOperationRequest) GetApiMethodName() string {
    return "yunos.trade.admin.common.operation"
}

func (r YunosTradeAdminCommonOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTradeAdminCommonOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r YunosTradeAdminCommonOperationRequest) GetParameter() string {
    return r.parameter
}

func (r *YunosTradeAdminCommonOperationRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

func (r YunosTradeAdminCommonOperationRequest) GetMethodName() string {
    return r.methodName
}

func (r *YunosTradeAdminCommonOperationRequest) SetInterfaceName(interfaceName string) error {
    r.interfaceName = interfaceName
    r.Set("interface_name", interfaceName)
    return nil
}

func (r YunosTradeAdminCommonOperationRequest) GetInterfaceName() string {
    return r.interfaceName
}

