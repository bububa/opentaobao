package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaceressupplierpoquerydetailAPIRequest 采购供应商订单明细查询接口 API请求
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
type AlibabaceressupplierpoquerydetailAPIRequest struct {
	model.Params
	// 订单编号
	_poNo string
}

// NewAlibabaceressupplierpoquerydetailRequest 初始化AlibabaceressupplierpoquerydetailAPIRequest对象
func NewAlibabaceressupplierpoquerydetailRequest() *AlibabaceressupplierpoquerydetailAPIRequest {
	return &AlibabaceressupplierpoquerydetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaceressupplierpoquerydetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaceressupplierpoquerydetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaceressupplierpoquerydetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoNo is PoNo Setter
// 订单编号
func (r *AlibabaceressupplierpoquerydetailAPIRequest) SetPoNo(_poNo string) error {
	r._poNo = _poNo
	r.Set("po_no", _poNo)
	return nil
}

// GetPoNo PoNo Getter
func (r AlibabaceressupplierpoquerydetailAPIRequest) GetPoNo() string {
	return r._poNo
}
