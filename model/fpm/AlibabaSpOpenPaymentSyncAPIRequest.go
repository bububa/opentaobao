package fpm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.sp.open.payment.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSpOpenPaymentSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
