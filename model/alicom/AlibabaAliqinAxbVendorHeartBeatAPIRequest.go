package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinaxbvendorheartbeatAPIRequest 供应商心跳上报接口 API请求
// alibaba.aliqin.axb.vendor.heart.beat
//
// 供应商上报自己的心跳信息
type AlibabaaliqinaxbvendorheartbeatAPIRequest struct {
	model.Params
	// 可选的预留字段
	_status string
	// 供应商合作KEY
	_vendorKey string
}

// NewAlibabaaliqinaxbvendorheartbeatRequest 初始化AlibabaaliqinaxbvendorheartbeatAPIRequest对象
func NewAlibabaaliqinaxbvendorheartbeatRequest() *AlibabaaliqinaxbvendorheartbeatAPIRequest {
	return &AlibabaaliqinaxbvendorheartbeatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinaxbvendorheartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.heart.beat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinaxbvendorheartbeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinaxbvendorheartbeatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 可选的预留字段
func (r *AlibabaaliqinaxbvendorheartbeatAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaaliqinaxbvendorheartbeatAPIRequest) GetStatus() string {
	return r._status
}

// SetVendorKey is VendorKey Setter
// 供应商合作KEY
func (r *AlibabaaliqinaxbvendorheartbeatAPIRequest) SetVendorKey(_vendorKey string) error {
	r._vendorKey = _vendorKey
	r.Set("vendor_key", _vendorKey)
	return nil
}

// GetVendorKey VendorKey Getter
func (r AlibabaaliqinaxbvendorheartbeatAPIRequest) GetVendorKey() string {
	return r._vendorKey
}
