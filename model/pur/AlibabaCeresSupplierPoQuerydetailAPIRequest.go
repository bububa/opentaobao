package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCeresSupplierPoQuerydetailAPIRequest 采购供应商订单明细查询接口 API请求
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
type AlibabaCeresSupplierPoQuerydetailAPIRequest struct {
	model.Params
	// 订单编号
	_poNo string
}

// NewAlibabaCeresSupplierPoQuerydetailRequest 初始化AlibabaCeresSupplierPoQuerydetailAPIRequest对象
func NewAlibabaCeresSupplierPoQuerydetailRequest() *AlibabaCeresSupplierPoQuerydetailAPIRequest {
	return &AlibabaCeresSupplierPoQuerydetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCeresSupplierPoQuerydetailAPIRequest) Reset() {
	r._poNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoNo is PoNo Setter
// 订单编号
func (r *AlibabaCeresSupplierPoQuerydetailAPIRequest) SetPoNo(_poNo string) error {
	r._poNo = _poNo
	r.Set("po_no", _poNo)
	return nil
}

// GetPoNo PoNo Getter
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetPoNo() string {
	return r._poNo
}

var poolAlibabaCeresSupplierPoQuerydetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCeresSupplierPoQuerydetailRequest()
	},
}

// GetAlibabaCeresSupplierPoQuerydetailRequest 从 sync.Pool 获取 AlibabaCeresSupplierPoQuerydetailAPIRequest
func GetAlibabaCeresSupplierPoQuerydetailAPIRequest() *AlibabaCeresSupplierPoQuerydetailAPIRequest {
	return poolAlibabaCeresSupplierPoQuerydetailAPIRequest.Get().(*AlibabaCeresSupplierPoQuerydetailAPIRequest)
}

// ReleaseAlibabaCeresSupplierPoQuerydetailAPIRequest 将 AlibabaCeresSupplierPoQuerydetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaCeresSupplierPoQuerydetailAPIRequest(v *AlibabaCeresSupplierPoQuerydetailAPIRequest) {
	v.Reset()
	poolAlibabaCeresSupplierPoQuerydetailAPIRequest.Put(v)
}
