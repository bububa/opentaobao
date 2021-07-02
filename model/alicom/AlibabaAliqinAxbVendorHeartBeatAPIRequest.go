package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorHeartBeatAPIRequest 供应商心跳上报接口 API请求
// alibaba.aliqin.axb.vendor.heart.beat
//
// 供应商上报自己的心跳信息
type AlibabaAliqinAxbVendorHeartBeatAPIRequest struct {
	model.Params
	// 可选的预留字段
	_status string
	// 供应商合作KEY
	_vendorKey string
}

// NewAlibabaAliqinAxbVendorHeartBeatRequest 初始化AlibabaAliqinAxbVendorHeartBeatAPIRequest对象
func NewAlibabaAliqinAxbVendorHeartBeatRequest() *AlibabaAliqinAxbVendorHeartBeatAPIRequest {
	return &AlibabaAliqinAxbVendorHeartBeatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.heart.beat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatus is Status Setter
// 可选的预留字段
func (r *AlibabaAliqinAxbVendorHeartBeatAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetStatus() string {
	return r._status
}

// SetVendorKey is VendorKey Setter
// 供应商合作KEY
func (r *AlibabaAliqinAxbVendorHeartBeatAPIRequest) SetVendorKey(_vendorKey string) error {
	r._vendorKey = _vendorKey
	r.Set("vendor_key", _vendorKey)
	return nil
}

// GetVendorKey VendorKey Getter
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetVendorKey() string {
	return r._vendorKey
}
