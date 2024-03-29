package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripInvoiceSettingAddAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.setting.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripInvoiceSettingAddAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripInvoiceSettingAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripInvoiceSettingAddRequest()
	},
}

// GetAlitripBtripInvoiceSettingAddRequest 从 sync.Pool 获取 AlitripBtripInvoiceSettingAddAPIRequest
func GetAlitripBtripInvoiceSettingAddAPIRequest() *AlitripBtripInvoiceSettingAddAPIRequest {
	return poolAlitripBtripInvoiceSettingAddAPIRequest.Get().(*AlitripBtripInvoiceSettingAddAPIRequest)
}

// ReleaseAlitripBtripInvoiceSettingAddAPIRequest 将 AlitripBtripInvoiceSettingAddAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripInvoiceSettingAddAPIRequest(v *AlitripBtripInvoiceSettingAddAPIRequest) {
	v.Reset()
	poolAlitripBtripInvoiceSettingAddAPIRequest.Put(v)
}
