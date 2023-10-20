package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderReceiptinfoGetAPIRequest Get Order Receipt Info API请求
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
type AliexpressSolutionOrderReceiptinfoGetAPIRequest struct {
	model.Params
	// query param
	_param1 *SingleOrderQuery
}

// NewAliexpressSolutionOrderReceiptinfoGetRequest 初始化AliexpressSolutionOrderReceiptinfoGetAPIRequest对象
func NewAliexpressSolutionOrderReceiptinfoGetRequest() *AliexpressSolutionOrderReceiptinfoGetAPIRequest {
	return &AliexpressSolutionOrderReceiptinfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionOrderReceiptinfoGetAPIRequest) Reset() {
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionOrderReceiptinfoGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.receiptinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionOrderReceiptinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionOrderReceiptinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// query param
func (r *AliexpressSolutionOrderReceiptinfoGetAPIRequest) SetParam1(_param1 *SingleOrderQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressSolutionOrderReceiptinfoGetAPIRequest) GetParam1() *SingleOrderQuery {
	return r._param1
}

var poolAliexpressSolutionOrderReceiptinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionOrderReceiptinfoGetRequest()
	},
}

// GetAliexpressSolutionOrderReceiptinfoGetRequest 从 sync.Pool 获取 AliexpressSolutionOrderReceiptinfoGetAPIRequest
func GetAliexpressSolutionOrderReceiptinfoGetAPIRequest() *AliexpressSolutionOrderReceiptinfoGetAPIRequest {
	return poolAliexpressSolutionOrderReceiptinfoGetAPIRequest.Get().(*AliexpressSolutionOrderReceiptinfoGetAPIRequest)
}

// ReleaseAliexpressSolutionOrderReceiptinfoGetAPIRequest 将 AliexpressSolutionOrderReceiptinfoGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionOrderReceiptinfoGetAPIRequest(v *AliexpressSolutionOrderReceiptinfoGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionOrderReceiptinfoGetAPIRequest.Put(v)
}
