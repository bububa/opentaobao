package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicesettingdeleteAPIRequest 发票删除 API请求
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
type AlitripbtripinvoicesettingdeleteAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenInvoiceDeleteRq
}

// NewAlitripbtripinvoicesettingdeleteRequest 初始化AlitripbtripinvoicesettingdeleteAPIRequest对象
func NewAlitripbtripinvoicesettingdeleteRequest() *AlitripbtripinvoicesettingdeleteAPIRequest {
	return &AlitripbtripinvoicesettingdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicesettingdeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicesettingdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicesettingdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripinvoicesettingdeleteAPIRequest) SetRq(_rq *OpenInvoiceDeleteRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripinvoicesettingdeleteAPIRequest) GetRq() *OpenInvoiceDeleteRq {
	return r._rq
}
