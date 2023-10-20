package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvmboscommonoperationAPIRequest 应用中心通用服务接口 API请求
// yunos.tvmbos.common.operation
//
// 应用中心相关接口的代理
type YunostvmboscommonoperationAPIRequest struct {
	model.Params
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
	// 入参，json格式
	_parameter string
}

// NewYunostvmboscommonoperationRequest 初始化YunostvmboscommonoperationAPIRequest对象
func NewYunostvmboscommonoperationRequest() *YunostvmboscommonoperationAPIRequest {
	return &YunostvmboscommonoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvmboscommonoperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvmbos.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvmboscommonoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvmboscommonoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名
func (r *YunostvmboscommonoperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunostvmboscommonoperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunostvmboscommonoperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunostvmboscommonoperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetParameter is Parameter Setter
// 入参，json格式
func (r *YunostvmboscommonoperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunostvmboscommonoperationAPIRequest) GetParameter() string {
	return r._parameter
}
