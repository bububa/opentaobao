package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingModifyAPIRequest 发票变更 API请求
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
type AlitripBtripInvoiceSettingModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// NewAlitripBtripInvoiceSettingModifyRequest 初始化AlitripBtripInvoiceSettingModifyAPIRequest对象
func NewAlitripBtripInvoiceSettingModifyRequest() *AlitripBtripInvoiceSettingModifyAPIRequest {
	return &AlitripBtripInvoiceSettingModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripInvoiceSettingModifyAPIRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetRq() *OpenInvoiceModifyAndNewRq {
	return r._rq
}
