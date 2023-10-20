package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillVerificateCallbackAPIRequest 翱象ERP核销回调 API请求
// alibaba.tcls.aelophy.bill.verificate.callback
//
// 翱象ERP核销回调
type AlibabaTclsAelophyBillVerificateCallbackAPIRequest struct {
	model.Params
	// 回调对象
	_module *VerificateCallbackDto
}

// NewAlibabaTclsAelophyBillVerificateCallbackRequest 初始化AlibabaTclsAelophyBillVerificateCallbackAPIRequest对象
func NewAlibabaTclsAelophyBillVerificateCallbackRequest() *AlibabaTclsAelophyBillVerificateCallbackAPIRequest {
	return &AlibabaTclsAelophyBillVerificateCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyBillVerificateCallbackAPIRequest) Reset() {
	r._module = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.verificate.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModule is Module Setter
// 回调对象
func (r *AlibabaTclsAelophyBillVerificateCallbackAPIRequest) SetModule(_module *VerificateCallbackDto) error {
	r._module = _module
	r.Set("module", _module)
	return nil
}

// GetModule Module Getter
func (r AlibabaTclsAelophyBillVerificateCallbackAPIRequest) GetModule() *VerificateCallbackDto {
	return r._module
}

var poolAlibabaTclsAelophyBillVerificateCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyBillVerificateCallbackRequest()
	},
}

// GetAlibabaTclsAelophyBillVerificateCallbackRequest 从 sync.Pool 获取 AlibabaTclsAelophyBillVerificateCallbackAPIRequest
func GetAlibabaTclsAelophyBillVerificateCallbackAPIRequest() *AlibabaTclsAelophyBillVerificateCallbackAPIRequest {
	return poolAlibabaTclsAelophyBillVerificateCallbackAPIRequest.Get().(*AlibabaTclsAelophyBillVerificateCallbackAPIRequest)
}

// ReleaseAlibabaTclsAelophyBillVerificateCallbackAPIRequest 将 AlibabaTclsAelophyBillVerificateCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyBillVerificateCallbackAPIRequest(v *AlibabaTclsAelophyBillVerificateCallbackAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyBillVerificateCallbackAPIRequest.Put(v)
}
