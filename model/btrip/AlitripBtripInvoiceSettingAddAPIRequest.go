package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingAddAPIRequest 发票设置 API请求
// alitrip.btrip.invoice.setting.add
//
// 发票设置
type AlitripBtripInvoiceSettingAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// NewAlitripBtripInvoiceSettingAddRequest 初始化AlitripBtripInvoiceSettingAddAPIRequest对象
func NewAlitripBtripInvoiceSettingAddRequest() *AlitripBtripInvoiceSettingAddAPIRequest {
	return &AlitripBtripInvoiceSettingAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripInvoiceSettingAddAPIRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetRq() *OpenInvoiceModifyAndNewRq {
	return r._rq
}
