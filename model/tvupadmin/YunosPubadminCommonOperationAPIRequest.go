package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunospubadmincommonoperationAPIRequest 内部迎客松通用服务 API请求
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
type YunospubadmincommonoperationAPIRequest struct {
	model.Params
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
	// 入参json串
	_parameter string
}

// NewYunospubadmincommonoperationRequest 初始化YunospubadmincommonoperationAPIRequest对象
func NewYunospubadmincommonoperationRequest() *YunospubadmincommonoperationAPIRequest {
	return &YunospubadmincommonoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunospubadmincommonoperationAPIRequest) GetApiMethodName() string {
	return "yunos.pubadmin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunospubadmincommonoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunospubadmincommonoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名
func (r *YunospubadmincommonoperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunospubadmincommonoperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunospubadmincommonoperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunospubadmincommonoperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetParameter is Parameter Setter
// 入参json串
func (r *YunospubadmincommonoperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunospubadmincommonoperationAPIRequest) GetParameter() string {
	return r._parameter
}
