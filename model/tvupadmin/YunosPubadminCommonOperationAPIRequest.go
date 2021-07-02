package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosPubadminCommonOperationAPIRequest 内部迎客松通用服务 API请求
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
type YunosPubadminCommonOperationAPIRequest struct {
	model.Params
	// 入参json串
	_parameter string
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
}

// NewYunosPubadminCommonOperationRequest 初始化YunosPubadminCommonOperationAPIRequest对象
func NewYunosPubadminCommonOperationRequest() *YunosPubadminCommonOperationAPIRequest {
	return &YunosPubadminCommonOperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosPubadminCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.pubadmin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosPubadminCommonOperationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParameter is Parameter Setter
// 入参json串
func (r *YunosPubadminCommonOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunosPubadminCommonOperationAPIRequest) GetParameter() string {
	return r._parameter
}

// SetInterfaceName is InterfaceName Setter
// 接口名
func (r *YunosPubadminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunosPubadminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunosPubadminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunosPubadminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}
