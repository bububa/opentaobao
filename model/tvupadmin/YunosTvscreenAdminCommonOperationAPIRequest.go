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
	// 接口名称
	_interfaceName string
	// 参数数组
	_parameters string
	// 方法名
	_methodName string
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
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名称
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetParameters is Parameters Setter
// 参数数组
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetParameters(_parameters string) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetParameters() string {
	return r._parameters
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunosTvscreenAdminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunosTvscreenAdminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}
