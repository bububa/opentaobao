package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvscreenadmincommonoperationAPIRequest 一体机桌面通用接口 API请求
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
type YunostvscreenadmincommonoperationAPIRequest struct {
	model.Params
	// 接口名称
	_interfaceName string
	// 参数数组
	_parameters string
	// 方法名
	_methodName string
}

// NewYunostvscreenadmincommonoperationRequest 初始化YunostvscreenadmincommonoperationAPIRequest对象
func NewYunostvscreenadmincommonoperationRequest() *YunostvscreenadmincommonoperationAPIRequest {
	return &YunostvscreenadmincommonoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvscreenadmincommonoperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvscreen.admin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvscreenadmincommonoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvscreenadmincommonoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名称
func (r *YunostvscreenadmincommonoperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunostvscreenadmincommonoperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetParameters is Parameters Setter
// 参数数组
func (r *YunostvscreenadmincommonoperationAPIRequest) SetParameters(_parameters string) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r YunostvscreenadmincommonoperationAPIRequest) GetParameters() string {
	return r._parameters
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunostvscreenadmincommonoperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunostvscreenadmincommonoperationAPIRequest) GetMethodName() string {
	return r._methodName
}
