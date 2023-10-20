package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingDeleteAPIRequest 发票删除 API请求
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
type AlitripBtripInvoiceSettingDeleteAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenInvoiceDeleteRq
}

// NewAlitripBtripInvoiceSettingDeleteRequest 初始化AlitripBtripInvoiceSettingDeleteAPIRequest对象
func NewAlitripBtripInvoiceSettingDeleteRequest() *AlitripBtripInvoiceSettingDeleteAPIRequest {
	return &AlitripBtripInvoiceSettingDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripInvoiceSettingDeleteAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripInvoiceSettingDeleteAPIRequest) SetRq(_rq *OpenInvoiceDeleteRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetRq() *OpenInvoiceDeleteRq {
	return r._rq
}

var poolAlitripBtripInvoiceSettingDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripInvoiceSettingDeleteRequest()
	},
}

// GetAlitripBtripInvoiceSettingDeleteRequest 从 sync.Pool 获取 AlitripBtripInvoiceSettingDeleteAPIRequest
func GetAlitripBtripInvoiceSettingDeleteAPIRequest() *AlitripBtripInvoiceSettingDeleteAPIRequest {
	return poolAlitripBtripInvoiceSettingDeleteAPIRequest.Get().(*AlitripBtripInvoiceSettingDeleteAPIRequest)
}

// ReleaseAlitripBtripInvoiceSettingDeleteAPIRequest 将 AlitripBtripInvoiceSettingDeleteAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripInvoiceSettingDeleteAPIRequest(v *AlitripBtripInvoiceSettingDeleteAPIRequest) {
	v.Reset()
	poolAlitripBtripInvoiceSettingDeleteAPIRequest.Put(v)
}
