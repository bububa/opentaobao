package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest
中心化供应商异常号码状态同步接口 API请求
alibaba.aliqin.axb.vendor.exception.no.sync

用于中心化供应商同步异常号码 */
type AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest struct {
	model.Params
	// 异常的中间号码
	_secretNo string
	// 异常的原因
	_exceptionMsg string
	// 0-异常状态 1-可恢复正常使用
	_status int64
	// 供应商KEY
	_vendorKey string
}

// NewAlibabaAliqinAxbVendorExceptionNoSyncRequest 初始化AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest对象
func NewAlibabaAliqinAxbVendorExceptionNoSyncRequest() *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest {
	return &AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.exception.no.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SecretNo Setter
// 异常的中间号码
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetSecretNo(_secretNo string) error {
	r._secretNo = _secretNo
	r.Set("secret_no", _secretNo)
	return nil
}

// Get SecretNo Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetSecretNo() string {
	return r._secretNo
}

// Set is ExceptionMsg Setter
// 异常的原因
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetExceptionMsg(_exceptionMsg string) error {
	r._exceptionMsg = _exceptionMsg
	r.Set("exception_msg", _exceptionMsg)
	return nil
}

// Get ExceptionMsg Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetExceptionMsg() string {
	return r._exceptionMsg
}

// Set is Status Setter
// 0-异常状态 1-可恢复正常使用
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is VendorKey Setter
// 供应商KEY
func (r *AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) SetVendorKey(_vendorKey string) error {
	r._vendorKey = _vendorKey
	r.Set("vendor_key", _vendorKey)
	return nil
}

// Get VendorKey Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest) GetVendorKey() string {
	return r._vendorKey
}
