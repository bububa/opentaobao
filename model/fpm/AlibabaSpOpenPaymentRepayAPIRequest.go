package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaspopenpaymentrepayAPIRequest 智付重新打款 API请求
// alibaba.sp.open.payment.repay
//
// 智付重新打款
type AlibabaspopenpaymentrepayAPIRequest struct {
	model.Params
	// 应用编码
	_appCode string
	// 入参数据
	_data string
}

// NewAlibabaspopenpaymentrepayRequest 初始化AlibabaspopenpaymentrepayAPIRequest对象
func NewAlibabaspopenpaymentrepayRequest() *AlibabaspopenpaymentrepayAPIRequest {
	return &AlibabaspopenpaymentrepayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaspopenpaymentrepayAPIRequest) GetApiMethodName() string {
	return "alibaba.sp.open.payment.repay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaspopenpaymentrepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaspopenpaymentrepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 应用编码
func (r *AlibabaspopenpaymentrepayAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaspopenpaymentrepayAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetData is Data Setter
// 入参数据
func (r *AlibabaspopenpaymentrepayAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaspopenpaymentrepayAPIRequest) GetData() string {
	return r._data
}
