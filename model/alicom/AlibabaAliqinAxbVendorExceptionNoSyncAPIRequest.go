package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest 中心化供应商异常号码状态同步接口 API请求
// alibaba.aliqin.axb.vendor.exception.no.sync
//
// 用于中心化供应商同步异常号码
type AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest struct {
	model.Params
	// 异常的中间号码
	_secretNo string
	// 异常的原因
	_exceptionMsg string
	// 供应商KEY
	_vendorKey string
	// 0-异常状态 1-可恢复正常使用
	_status int64
}

// NewAlibabaAliqinAxbVendorExceptionNoSyncRequest 初始化AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest对象
func NewAlibabaAliqinAxbVendorExceptionNoSyncRequest() *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest {
	return &AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) Reset() {
	r._secretNo = ""
	r._exceptionMsg = ""
	r._vendorKey = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.exception.no.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecretNo is SecretNo Setter
// 异常的中间号码
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetSecretNo(_secretNo string) error {
	r._secretNo = _secretNo
	r.Set("secret_no", _secretNo)
	return nil
}

// GetSecretNo SecretNo Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetSecretNo() string {
	return r._secretNo
}

// SetExceptionMsg is ExceptionMsg Setter
// 异常的原因
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetExceptionMsg(_exceptionMsg string) error {
	r._exceptionMsg = _exceptionMsg
	r.Set("exception_msg", _exceptionMsg)
	return nil
}

// GetExceptionMsg ExceptionMsg Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetExceptionMsg() string {
	return r._exceptionMsg
}

// SetVendorKey is VendorKey Setter
// 供应商KEY
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetVendorKey(_vendorKey string) error {
	r._vendorKey = _vendorKey
	r.Set("vendor_key", _vendorKey)
	return nil
}

// GetVendorKey VendorKey Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetVendorKey() string {
	return r._vendorKey
}

// SetStatus is Status Setter
// 0-异常状态 1-可恢复正常使用
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinAxbVendorExceptionNoSyncRequest()
	},
}

// GetAlibabaAliqinAxbVendorExceptionNoSyncRequest 从 sync.Pool 获取 AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest
func GetAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest() *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest {
	return poolAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest.Get().(*AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest)
}

// ReleaseAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest 将 AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest(v *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) {
	v.Reset()
	poolAlibabaAliqinAxbVendorExceptionNoSyncAPIRequest.Put(v)
}
