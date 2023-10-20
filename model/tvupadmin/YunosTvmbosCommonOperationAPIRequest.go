package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvmbosCommonOperationAPIRequest 应用中心通用服务接口 API请求
// yunos.tvmbos.common.operation
//
// 应用中心相关接口的代理
type YunosTvmbosCommonOperationAPIRequest struct {
	model.Params
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
	// 入参，json格式
	_parameter string
}

// NewYunosTvmbosCommonOperationRequest 初始化YunosTvmbosCommonOperationAPIRequest对象
func NewYunosTvmbosCommonOperationRequest() *YunosTvmbosCommonOperationAPIRequest {
	return &YunosTvmbosCommonOperationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvmbosCommonOperationAPIRequest) Reset() {
	r._interfaceName = ""
	r._methodName = ""
	r._parameter = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvmbosCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvmbos.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvmbosCommonOperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvmbosCommonOperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名
func (r *YunosTvmbosCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunosTvmbosCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunosTvmbosCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunosTvmbosCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetParameter is Parameter Setter
// 入参，json格式
func (r *YunosTvmbosCommonOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunosTvmbosCommonOperationAPIRequest) GetParameter() string {
	return r._parameter
}

var poolYunosTvmbosCommonOperationAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvmbosCommonOperationRequest()
	},
}

// GetYunosTvmbosCommonOperationRequest 从 sync.Pool 获取 YunosTvmbosCommonOperationAPIRequest
func GetYunosTvmbosCommonOperationAPIRequest() *YunosTvmbosCommonOperationAPIRequest {
	return poolYunosTvmbosCommonOperationAPIRequest.Get().(*YunosTvmbosCommonOperationAPIRequest)
}

// ReleaseYunosTvmbosCommonOperationAPIRequest 将 YunosTvmbosCommonOperationAPIRequest 放入 sync.Pool
func ReleaseYunosTvmbosCommonOperationAPIRequest(v *YunosTvmbosCommonOperationAPIRequest) {
	v.Reset()
	poolYunosTvmbosCommonOperationAPIRequest.Put(v)
}
