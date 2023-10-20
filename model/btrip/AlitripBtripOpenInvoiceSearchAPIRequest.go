package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopeninvoicesearchAPIRequest 差旅申请用户搜索可用发票列表 API请求
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
type AlitripbtripopeninvoicesearchAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRq
}

// NewAlitripbtripopeninvoicesearchRequest 初始化AlitripbtripopeninvoicesearchAPIRequest对象
func NewAlitripbtripopeninvoicesearchRequest() *AlitripbtripopeninvoicesearchAPIRequest {
	return &AlitripbtripopeninvoicesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopeninvoicesearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.invoice.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopeninvoicesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopeninvoicesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopeninvoicesearchAPIRequest) SetRq(_rq *OpenInvoiceRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopeninvoicesearchAPIRequest) GetRq() *OpenInvoiceRq {
	return r._rq
}
