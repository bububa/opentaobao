package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpPresalespackageConsignAPIRequest 预售预包尾款推单发货 API请求
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
type AlibabaAscpPresalespackageConsignAPIRequest struct {
	model.Params
	// 入参
	_requestParams *Requestparams
}

// NewAlibabaAscpPresalespackageConsignRequest 初始化AlibabaAscpPresalespackageConsignAPIRequest对象
func NewAlibabaAscpPresalespackageConsignRequest() *AlibabaAscpPresalespackageConsignAPIRequest {
	return &AlibabaAscpPresalespackageConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.presalespackage.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequestParams is RequestParams Setter
// 入参
func (r *AlibabaAscpPresalespackageConsignAPIRequest) SetRequestParams(_requestParams *Requestparams) error {
	r._requestParams = _requestParams
	r.Set("request_params", _requestParams)
	return nil
}

// GetRequestParams RequestParams Getter
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetRequestParams() *Requestparams {
	return r._requestParams
}
