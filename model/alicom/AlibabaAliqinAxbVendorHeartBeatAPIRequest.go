package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinAxbVendorHeartBeatAPIRequest) Reset() {
	r._status = ""
	r._vendorKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.heart.beat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinAxbVendorHeartBeatAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAliqinAxbVendorHeartBeatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinAxbVendorHeartBeatRequest()
	},
}

// GetAlibabaAliqinAxbVendorHeartBeatRequest 从 sync.Pool 获取 AlibabaAliqinAxbVendorHeartBeatAPIRequest
func GetAlibabaAliqinAxbVendorHeartBeatAPIRequest() *AlibabaAliqinAxbVendorHeartBeatAPIRequest {
	return poolAlibabaAliqinAxbVendorHeartBeatAPIRequest.Get().(*AlibabaAliqinAxbVendorHeartBeatAPIRequest)
}

// ReleaseAlibabaAliqinAxbVendorHeartBeatAPIRequest 将 AlibabaAliqinAxbVendorHeartBeatAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinAxbVendorHeartBeatAPIRequest(v *AlibabaAliqinAxbVendorHeartBeatAPIRequest) {
	v.Reset()
	poolAlibabaAliqinAxbVendorHeartBeatAPIRequest.Put(v)
}
