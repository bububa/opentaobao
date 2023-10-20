package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicesettingmodifyAPIRequest 发票变更 API请求
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
type AlitripbtripinvoicesettingmodifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// NewAlitripbtripinvoicesettingmodifyRequest 初始化AlitripbtripinvoicesettingmodifyAPIRequest对象
func NewAlitripbtripinvoicesettingmodifyRequest() *AlitripbtripinvoicesettingmodifyAPIRequest {
	return &AlitripbtripinvoicesettingmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicesettingmodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicesettingmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicesettingmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripinvoicesettingmodifyAPIRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripinvoicesettingmodifyAPIRequest) GetRq() *OpenInvoiceModifyAndNewRq {
	return r._rq
}
