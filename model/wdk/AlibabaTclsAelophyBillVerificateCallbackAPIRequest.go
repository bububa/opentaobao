package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophybillverificatecallbackAPIRequest 翱象ERP核销回调 API请求
// alibaba.tcls.aelophy.bill.verificate.callback
//
// 翱象ERP核销回调
type AlibabatclsaelophybillverificatecallbackAPIRequest struct {
	model.Params
	// 回调对象
	_module *VerificateCallbackDto
}

// NewAlibabatclsaelophybillverificatecallbackRequest 初始化AlibabatclsaelophybillverificatecallbackAPIRequest对象
func NewAlibabatclsaelophybillverificatecallbackRequest() *AlibabatclsaelophybillverificatecallbackAPIRequest {
	return &AlibabatclsaelophybillverificatecallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophybillverificatecallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.verificate.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophybillverificatecallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophybillverificatecallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModule is Module Setter
// 回调对象
func (r *AlibabatclsaelophybillverificatecallbackAPIRequest) SetModule(_module *VerificateCallbackDto) error {
	r._module = _module
	r.Set("module", _module)
	return nil
}

// GetModule Module Getter
func (r AlibabatclsaelophybillverificatecallbackAPIRequest) GetModule() *VerificateCallbackDto {
	return r._module
}
