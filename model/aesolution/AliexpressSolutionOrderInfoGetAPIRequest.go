package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderInfoGetAPIRequest get order detail info API请求
// aliexpress.solution.order.info.get
//
// get order detail info
type AliexpressSolutionOrderInfoGetAPIRequest struct {
	model.Params
	// param
	_param1 *OrderDetailQuery
}

// NewAliexpressSolutionOrderInfoGetRequest 初始化AliexpressSolutionOrderInfoGetAPIRequest对象
func NewAliexpressSolutionOrderInfoGetRequest() *AliexpressSolutionOrderInfoGetAPIRequest {
	return &AliexpressSolutionOrderInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionOrderInfoGetAPIRequest) Reset() {
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpressSolutionOrderInfoGetAPIRequest) SetParam1(_param1 *OrderDetailQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressSolutionOrderInfoGetAPIRequest) GetParam1() *OrderDetailQuery {
	return r._param1
}

var poolAliexpressSolutionOrderInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionOrderInfoGetRequest()
	},
}

// GetAliexpressSolutionOrderInfoGetRequest 从 sync.Pool 获取 AliexpressSolutionOrderInfoGetAPIRequest
func GetAliexpressSolutionOrderInfoGetAPIRequest() *AliexpressSolutionOrderInfoGetAPIRequest {
	return poolAliexpressSolutionOrderInfoGetAPIRequest.Get().(*AliexpressSolutionOrderInfoGetAPIRequest)
}

// ReleaseAliexpressSolutionOrderInfoGetAPIRequest 将 AliexpressSolutionOrderInfoGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionOrderInfoGetAPIRequest(v *AliexpressSolutionOrderInfoGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionOrderInfoGetAPIRequest.Put(v)
}
