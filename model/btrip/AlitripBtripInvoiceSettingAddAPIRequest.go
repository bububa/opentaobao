package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicesettingaddAPIRequest 发票设置 API请求
// alitrip.btrip.invoice.setting.add
//
// 发票设置
type AlitripbtripinvoicesettingaddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// NewAlitripbtripinvoicesettingaddRequest 初始化AlitripbtripinvoicesettingaddAPIRequest对象
func NewAlitripbtripinvoicesettingaddRequest() *AlitripbtripinvoicesettingaddAPIRequest {
	return &AlitripbtripinvoicesettingaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicesettingaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicesettingaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicesettingaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripinvoicesettingaddAPIRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripinvoicesettingaddAPIRequest) GetRq() *OpenInvoiceModifyAndNewRq {
	return r._rq
}
