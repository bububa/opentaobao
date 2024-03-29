package fpm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSpOpenPaymentSyncAPIRequest 付款单同步 API请求
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
type AlibabaSpOpenPaymentSyncAPIRequest struct {
	model.Params
	// 应用编码
	_appCode string
	// 入参数据
	_data string
}

// NewAlibabaSpOpenPaymentSyncRequest 初始化AlibabaSpOpenPaymentSyncAPIRequest对象
func NewAlibabaSpOpenPaymentSyncRequest() *AlibabaSpOpenPaymentSyncAPIRequest {
	return &AlibabaSpOpenPaymentSyncAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSpOpenPaymentSyncAPIRequest) Reset() {
	r._appCode = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.sp.open.payment.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 应用编码
func (r *AlibabaSpOpenPaymentSyncAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetData is Data Setter
// 入参数据
func (r *AlibabaSpOpenPaymentSyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetData() string {
	return r._data
}

var poolAlibabaSpOpenPaymentSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSpOpenPaymentSyncRequest()
	},
}

// GetAlibabaSpOpenPaymentSyncRequest 从 sync.Pool 获取 AlibabaSpOpenPaymentSyncAPIRequest
func GetAlibabaSpOpenPaymentSyncAPIRequest() *AlibabaSpOpenPaymentSyncAPIRequest {
	return poolAlibabaSpOpenPaymentSyncAPIRequest.Get().(*AlibabaSpOpenPaymentSyncAPIRequest)
}

// ReleaseAlibabaSpOpenPaymentSyncAPIRequest 将 AlibabaSpOpenPaymentSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaSpOpenPaymentSyncAPIRequest(v *AlibabaSpOpenPaymentSyncAPIRequest) {
	v.Reset()
	poolAlibabaSpOpenPaymentSyncAPIRequest.Put(v)
}
