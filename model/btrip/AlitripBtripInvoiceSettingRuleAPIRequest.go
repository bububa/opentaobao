package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicesettingruleAPIRequest 发票规则设置 API请求
// alitrip.btrip.invoice.setting.rule
//
// 发票规则设置
type AlitripbtripinvoicesettingruleAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRuleRq
}

// NewAlitripbtripinvoicesettingruleRequest 初始化AlitripbtripinvoicesettingruleAPIRequest对象
func NewAlitripbtripinvoicesettingruleRequest() *AlitripbtripinvoicesettingruleAPIRequest {
	return &AlitripbtripinvoicesettingruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicesettingruleAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.rule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicesettingruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicesettingruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripinvoicesettingruleAPIRequest) SetRq(_rq *OpenInvoiceRuleRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripinvoicesettingruleAPIRequest) GetRq() *OpenInvoiceRuleRq {
	return r._rq
}
