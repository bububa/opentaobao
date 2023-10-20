package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripInvoiceSettingModifyAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripInvoiceSettingModifyAPIRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripInvoiceSettingModifyAPIRequest) GetRq() *OpenInvoiceModifyAndNewRq {
	return r._rq
}

var poolAlitripBtripInvoiceSettingModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripInvoiceSettingModifyRequest()
	},
}

// GetAlitripBtripInvoiceSettingModifyRequest 从 sync.Pool 获取 AlitripBtripInvoiceSettingModifyAPIRequest
func GetAlitripBtripInvoiceSettingModifyAPIRequest() *AlitripBtripInvoiceSettingModifyAPIRequest {
	return poolAlitripBtripInvoiceSettingModifyAPIRequest.Get().(*AlitripBtripInvoiceSettingModifyAPIRequest)
}

// ReleaseAlitripBtripInvoiceSettingModifyAPIRequest 将 AlitripBtripInvoiceSettingModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripInvoiceSettingModifyAPIRequest(v *AlitripBtripInvoiceSettingModifyAPIRequest) {
	v.Reset()
	poolAlitripBtripInvoiceSettingModifyAPIRequest.Put(v)
}
