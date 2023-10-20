package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascppresalespackageconsignAPIRequest 预售预包尾款推单发货 API请求
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
type AlibabaascppresalespackageconsignAPIRequest struct {
	model.Params
	// 入参
	_requestParams *Requestparams
}

// NewAlibabaascppresalespackageconsignRequest 初始化AlibabaascppresalespackageconsignAPIRequest对象
func NewAlibabaascppresalespackageconsignRequest() *AlibabaascppresalespackageconsignAPIRequest {
	return &AlibabaascppresalespackageconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascppresalespackageconsignAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.presalespackage.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascppresalespackageconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascppresalespackageconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParams is RequestParams Setter
// 入参
func (r *AlibabaascppresalespackageconsignAPIRequest) SetRequestParams(_requestParams *Requestparams) error {
	r._requestParams = _requestParams
	r.Set("request_params", _requestParams)
	return nil
}

// GetRequestParams RequestParams Getter
func (r AlibabaascppresalespackageconsignAPIRequest) GetRequestParams() *Requestparams {
	return r._requestParams
}
