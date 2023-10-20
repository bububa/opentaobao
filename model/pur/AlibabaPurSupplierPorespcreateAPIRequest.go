package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierPorespcreateAPIRequest po反馈创建 API请求
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
type AlibabaPurSupplierPorespcreateAPIRequest struct {
	model.Params
	// PO反馈信息
	_poResponse []SupplierPoResponseDo
}

// NewAlibabaPurSupplierPorespcreateRequest 初始化AlibabaPurSupplierPorespcreateAPIRequest对象
func NewAlibabaPurSupplierPorespcreateRequest() *AlibabaPurSupplierPorespcreateAPIRequest {
	return &AlibabaPurSupplierPorespcreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurSupplierPorespcreateAPIRequest) Reset() {
	r._poResponse = r._poResponse[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierPorespcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.porespcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierPorespcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurSupplierPorespcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoResponse is PoResponse Setter
// PO反馈信息
func (r *AlibabaPurSupplierPorespcreateAPIRequest) SetPoResponse(_poResponse []SupplierPoResponseDo) error {
	r._poResponse = _poResponse
	r.Set("po_response", _poResponse)
	return nil
}

// GetPoResponse PoResponse Getter
func (r AlibabaPurSupplierPorespcreateAPIRequest) GetPoResponse() []SupplierPoResponseDo {
	return r._poResponse
}

var poolAlibabaPurSupplierPorespcreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurSupplierPorespcreateRequest()
	},
}

// GetAlibabaPurSupplierPorespcreateRequest 从 sync.Pool 获取 AlibabaPurSupplierPorespcreateAPIRequest
func GetAlibabaPurSupplierPorespcreateAPIRequest() *AlibabaPurSupplierPorespcreateAPIRequest {
	return poolAlibabaPurSupplierPorespcreateAPIRequest.Get().(*AlibabaPurSupplierPorespcreateAPIRequest)
}

// ReleaseAlibabaPurSupplierPorespcreateAPIRequest 将 AlibabaPurSupplierPorespcreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurSupplierPorespcreateAPIRequest(v *AlibabaPurSupplierPorespcreateAPIRequest) {
	v.Reset()
	poolAlibabaPurSupplierPorespcreateAPIRequest.Put(v)
}
