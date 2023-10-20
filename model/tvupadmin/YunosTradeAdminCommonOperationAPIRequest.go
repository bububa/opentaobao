package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTradeAdminCommonOperationAPIRequest 交易迎客松通用服务接口 API请求
// yunos.trade.admin.common.operation
//
// 迎客松交易相关通用接口
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTradeAdminCommonOperationAPIRequest) Reset() {
	r._parameter = ""
	r._methodName = ""
	r._interfaceName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTradeAdminCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.trade.admin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTradeAdminCommonOperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTradeAdminCommonOperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 入参数据，json格式
func (r *YunosTradeAdminCommonOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetParameter() string {
	return r._parameter
}

// SetMethodName is MethodName Setter
// 调用方法
func (r *YunosTradeAdminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetInterfaceName is InterfaceName Setter
// 调用接口
func (r *YunosTradeAdminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunosTradeAdminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

var poolYunosTradeAdminCommonOperationAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTradeAdminCommonOperationRequest()
	},
}

// GetYunosTradeAdminCommonOperationRequest 从 sync.Pool 获取 YunosTradeAdminCommonOperationAPIRequest
func GetYunosTradeAdminCommonOperationAPIRequest() *YunosTradeAdminCommonOperationAPIRequest {
	return poolYunosTradeAdminCommonOperationAPIRequest.Get().(*YunosTradeAdminCommonOperationAPIRequest)
}

// ReleaseYunosTradeAdminCommonOperationAPIRequest 将 YunosTradeAdminCommonOperationAPIRequest 放入 sync.Pool
func ReleaseYunosTradeAdminCommonOperationAPIRequest(v *YunosTradeAdminCommonOperationAPIRequest) {
	v.Reset()
	poolYunosTradeAdminCommonOperationAPIRequest.Put(v)
}
