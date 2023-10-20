package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillDetailQueryAPIRequest 账单明细接口 API请求
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
type AlibabaTclsAelophyBillDetailQueryAPIRequest struct {
	model.Params
	// 请求对象
	_detailRequest *BillDetailQueryRequest
}

// NewAlibabaTclsAelophyBillDetailQueryRequest 初始化AlibabaTclsAelophyBillDetailQueryAPIRequest对象
func NewAlibabaTclsAelophyBillDetailQueryRequest() *AlibabaTclsAelophyBillDetailQueryAPIRequest {
	return &AlibabaTclsAelophyBillDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyBillDetailQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyBillDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyBillDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailRequest is DetailRequest Setter
// 请求对象
func (r *AlibabaTclsAelophyBillDetailQueryAPIRequest) SetDetailRequest(_detailRequest *BillDetailQueryRequest) error {
	r._detailRequest = _detailRequest
	r.Set("detail_request", _detailRequest)
	return nil
}

// GetDetailRequest DetailRequest Getter
func (r AlibabaTclsAelophyBillDetailQueryAPIRequest) GetDetailRequest() *BillDetailQueryRequest {
	return r._detailRequest
}
