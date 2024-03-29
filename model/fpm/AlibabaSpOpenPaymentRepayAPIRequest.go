package fpm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSpOpenPaymentRepayAPIRequest 智付重新打款 API请求
// alibaba.sp.open.payment.repay
//
// 智付重新打款
type AlibabaSpOpenPaymentRepayAPIRequest struct {
	model.Params
	// 应用编码
	_appCode string
	// 入参数据
	_data string
}

// NewAlibabaSpOpenPaymentRepayRequest 初始化AlibabaSpOpenPaymentRepayAPIRequest对象
func NewAlibabaSpOpenPaymentRepayRequest() *AlibabaSpOpenPaymentRepayAPIRequest {
	return &AlibabaSpOpenPaymentRepayAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSpOpenPaymentRepayAPIRequest) Reset() {
	r._appCode = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSpOpenPaymentRepayAPIRequest) GetApiMethodName() string {
	return "alibaba.sp.open.payment.repay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSpOpenPaymentRepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSpOpenPaymentRepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 应用编码
func (r *AlibabaSpOpenPaymentRepayAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaSpOpenPaymentRepayAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetData is Data Setter
// 入参数据
func (r *AlibabaSpOpenPaymentRepayAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaSpOpenPaymentRepayAPIRequest) GetData() string {
	return r._data
}

var poolAlibabaSpOpenPaymentRepayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSpOpenPaymentRepayRequest()
	},
}

// GetAlibabaSpOpenPaymentRepayRequest 从 sync.Pool 获取 AlibabaSpOpenPaymentRepayAPIRequest
func GetAlibabaSpOpenPaymentRepayAPIRequest() *AlibabaSpOpenPaymentRepayAPIRequest {
	return poolAlibabaSpOpenPaymentRepayAPIRequest.Get().(*AlibabaSpOpenPaymentRepayAPIRequest)
}

// ReleaseAlibabaSpOpenPaymentRepayAPIRequest 将 AlibabaSpOpenPaymentRepayAPIRequest 放入 sync.Pool
func ReleaseAlibabaSpOpenPaymentRepayAPIRequest(v *AlibabaSpOpenPaymentRepayAPIRequest) {
	v.Reset()
	poolAlibabaSpOpenPaymentRepayAPIRequest.Put(v)
}
