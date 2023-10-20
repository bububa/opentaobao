package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapursupplierporespcreateAPIRequest po反馈创建 API请求
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
type AlibabapursupplierporespcreateAPIRequest struct {
	model.Params
	// PO反馈信息
	_poResponse []SupplierPoResponseDo
}

// NewAlibabapursupplierporespcreateRequest 初始化AlibabapursupplierporespcreateAPIRequest对象
func NewAlibabapursupplierporespcreateRequest() *AlibabapursupplierporespcreateAPIRequest {
	return &AlibabapursupplierporespcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapursupplierporespcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.porespcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapursupplierporespcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapursupplierporespcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoResponse is PoResponse Setter
// PO反馈信息
func (r *AlibabapursupplierporespcreateAPIRequest) SetPoResponse(_poResponse []SupplierPoResponseDo) error {
	r._poResponse = _poResponse
	r.Set("po_response", _poResponse)
	return nil
}

// GetPoResponse PoResponse Getter
func (r AlibabapursupplierporespcreateAPIRequest) GetPoResponse() []SupplierPoResponseDo {
	return r._poResponse
}
