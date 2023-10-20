package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostradeadmincommonoperationAPIRequest 交易迎客松通用服务接口 API请求
// yunos.trade.admin.common.operation
//
// 迎客松交易相关通用接口
type YunostradeadmincommonoperationAPIRequest struct {
	model.Params
	// 入参数据，json格式
	_parameter string
	// 调用方法
	_methodName string
	// 调用接口
	_interfaceName string
}

// NewYunostradeadmincommonoperationRequest 初始化YunostradeadmincommonoperationAPIRequest对象
func NewYunostradeadmincommonoperationRequest() *YunostradeadmincommonoperationAPIRequest {
	return &YunostradeadmincommonoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostradeadmincommonoperationAPIRequest) GetApiMethodName() string {
	return "yunos.trade.admin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostradeadmincommonoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostradeadmincommonoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 入参数据，json格式
func (r *YunostradeadmincommonoperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunostradeadmincommonoperationAPIRequest) GetParameter() string {
	return r._parameter
}

// SetMethodName is MethodName Setter
// 调用方法
func (r *YunostradeadmincommonoperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunostradeadmincommonoperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetInterfaceName is InterfaceName Setter
// 调用接口
func (r *YunostradeadmincommonoperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunostradeadmincommonoperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}
