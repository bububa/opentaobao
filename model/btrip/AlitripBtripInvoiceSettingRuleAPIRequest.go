package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingRuleAPIRequest 发票规则设置 API请求
// alitrip.btrip.invoice.setting.rule
//
// 发票规则设置
type AlitripBtripInvoiceSettingRuleAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRuleRq
}

// NewAlitripBtripInvoiceSettingRuleRequest 初始化AlitripBtripInvoiceSettingRuleAPIRequest对象
func NewAlitripBtripInvoiceSettingRuleRequest() *AlitripBtripInvoiceSettingRuleAPIRequest {
	return &AlitripBtripInvoiceSettingRuleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripInvoiceSettingRuleAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.rule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripInvoiceSettingRuleAPIRequest) SetRq(_rq *OpenInvoiceRuleRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetRq() *OpenInvoiceRuleRq {
	return r._rq
}

var poolAlitripBtripInvoiceSettingRuleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripInvoiceSettingRuleRequest()
	},
}

// GetAlitripBtripInvoiceSettingRuleRequest 从 sync.Pool 获取 AlitripBtripInvoiceSettingRuleAPIRequest
func GetAlitripBtripInvoiceSettingRuleAPIRequest() *AlitripBtripInvoiceSettingRuleAPIRequest {
	return poolAlitripBtripInvoiceSettingRuleAPIRequest.Get().(*AlitripBtripInvoiceSettingRuleAPIRequest)
}

// ReleaseAlitripBtripInvoiceSettingRuleAPIRequest 将 AlitripBtripInvoiceSettingRuleAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripInvoiceSettingRuleAPIRequest(v *AlitripBtripInvoiceSettingRuleAPIRequest) {
	v.Reset()
	poolAlitripBtripInvoiceSettingRuleAPIRequest.Put(v)
}
