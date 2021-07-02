package pur

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQuerydetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
