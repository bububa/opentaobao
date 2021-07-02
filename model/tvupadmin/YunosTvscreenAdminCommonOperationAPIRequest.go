package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvscreenAdminCommonOperationAPIRequest 一体机桌面通用接口 API请求
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
type YunosTvscreenAdminCommonOperationAPIRequest struct {
	model.Params
	// 参数数组
	_parameters string
	// 方法名
	_methodName string
	// 接口名称
	_interfaceName string
}

// NewYunosTvscreenAdminCommonOperationRequest 初始化YunosTvscreenAdminCommonOperationAPIRequest对象
func NewYunosTvscreenAdminCommonOperationRequest() *YunosTvscreenAdminCommonOperationAPIRequest {
	return &YunosTvscreenAdminCommonOperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvscreen.admin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Parameters Setter
// 参数数组
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetParameters(_parameters string) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// Get Parameters Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetParameters() string {
	return r._parameters
}

// Set is MethodName Setter
// 方法名
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// Get MethodName Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// Set is InterfaceName Setter
// 接口名称
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// Get InterfaceName Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}
