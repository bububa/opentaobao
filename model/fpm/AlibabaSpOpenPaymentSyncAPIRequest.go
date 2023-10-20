package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaspopenpaymentsyncAPIRequest 付款单同步 API请求
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
type AlibabaspopenpaymentsyncAPIRequest struct {
	model.Params
	// 应用编码
	_appCode string
	// 入参数据
	_data string
}

// NewAlibabaspopenpaymentsyncRequest 初始化AlibabaspopenpaymentsyncAPIRequest对象
func NewAlibabaspopenpaymentsyncRequest() *AlibabaspopenpaymentsyncAPIRequest {
	return &AlibabaspopenpaymentsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaspopenpaymentsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.sp.open.payment.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaspopenpaymentsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaspopenpaymentsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 应用编码
func (r *AlibabaspopenpaymentsyncAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaspopenpaymentsyncAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetData is Data Setter
// 入参数据
func (r *AlibabaspopenpaymentsyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaspopenpaymentsyncAPIRequest) GetData() string {
	return r._data
}
