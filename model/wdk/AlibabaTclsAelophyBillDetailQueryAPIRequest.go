package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophybilldetailqueryAPIRequest 账单明细接口 API请求
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
type AlibabatclsaelophybilldetailqueryAPIRequest struct {
	model.Params
	// 请求对象
	_detailRequest *BillDetailQueryRequest
}

// NewAlibabatclsaelophybilldetailqueryRequest 初始化AlibabatclsaelophybilldetailqueryAPIRequest对象
func NewAlibabatclsaelophybilldetailqueryRequest() *AlibabatclsaelophybilldetailqueryAPIRequest {
	return &AlibabatclsaelophybilldetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophybilldetailqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophybilldetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophybilldetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailRequest is DetailRequest Setter
// 请求对象
func (r *AlibabatclsaelophybilldetailqueryAPIRequest) SetDetailRequest(_detailRequest *BillDetailQueryRequest) error {
	r._detailRequest = _detailRequest
	r.Set("detail_request", _detailRequest)
	return nil
}

// GetDetailRequest DetailRequest Getter
func (r AlibabatclsaelophybilldetailqueryAPIRequest) GetDetailRequest() *BillDetailQueryRequest {
	return r._detailRequest
}
