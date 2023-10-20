package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderGetAPIRequest get order list API请求
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
type AliexpressSolutionOrderGetAPIRequest struct {
	model.Params
	// param
	_param0 *OrderQuery
}

// NewAliexpressSolutionOrderGetRequest 初始化AliexpressSolutionOrderGetAPIRequest对象
func NewAliexpressSolutionOrderGetRequest() *AliexpressSolutionOrderGetAPIRequest {
	return &AliexpressSolutionOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionOrderGetAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// param
func (r *AliexpressSolutionOrderGetAPIRequest) SetParam0(_param0 *OrderQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AliexpressSolutionOrderGetAPIRequest) GetParam0() *OrderQuery {
	return r._param0
}

var poolAliexpressSolutionOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionOrderGetRequest()
	},
}

// GetAliexpressSolutionOrderGetRequest 从 sync.Pool 获取 AliexpressSolutionOrderGetAPIRequest
func GetAliexpressSolutionOrderGetAPIRequest() *AliexpressSolutionOrderGetAPIRequest {
	return poolAliexpressSolutionOrderGetAPIRequest.Get().(*AliexpressSolutionOrderGetAPIRequest)
}

// ReleaseAliexpressSolutionOrderGetAPIRequest 将 AliexpressSolutionOrderGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionOrderGetAPIRequest(v *AliexpressSolutionOrderGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionOrderGetAPIRequest.Put(v)
}
