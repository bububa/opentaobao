package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTradeAdminCommonOperationAPIRequest
交易迎客松通用服务接口 API请求
yunos.trade.admin.common.operation

迎客松交易相关通用接口 */
type YunosTradeAdminCommonOperationAPIRequest struct {
	model.Params
	// 入参数据，json格式
	_parameter string
	// 调用方法
	_methodName string
	// 调用接口
	_interfaceName string
}

// NewYunosTradeAdminCommonOperationRequest 初始化YunosTradeAdminCommonOperationAPIRequest对象
func NewYunosTradeAdminCommonOperationRequest() *YunosTradeAdminCommonOperationAPIRequest {
	return &YunosTradeAdminCommonOperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTradeAdminCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.trade.admin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTradeAdminCommonOperationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Parameter Setter
// 入参数据，json格式
func (r *YunosTradeAdminCommonOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// Get Parameter Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetParameter() string {
	return r._parameter
}

// Set is MethodName Setter
// 调用方法
func (r *YunosTradeAdminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// Get MethodName Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// Set is InterfaceName Setter
// 调用接口
func (r *YunosTradeAdminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// Get InterfaceName Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}
