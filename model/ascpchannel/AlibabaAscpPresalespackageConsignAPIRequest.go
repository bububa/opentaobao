package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpPresalespackageConsignAPIRequest) Reset() {
	r._requestParams = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.presalespackage.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpPresalespackageConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpPresalespackageConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpPresalespackageConsignRequest()
	},
}

// GetAlibabaAscpPresalespackageConsignRequest 从 sync.Pool 获取 AlibabaAscpPresalespackageConsignAPIRequest
func GetAlibabaAscpPresalespackageConsignAPIRequest() *AlibabaAscpPresalespackageConsignAPIRequest {
	return poolAlibabaAscpPresalespackageConsignAPIRequest.Get().(*AlibabaAscpPresalespackageConsignAPIRequest)
}

// ReleaseAlibabaAscpPresalespackageConsignAPIRequest 将 AlibabaAscpPresalespackageConsignAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpPresalespackageConsignAPIRequest(v *AlibabaAscpPresalespackageConsignAPIRequest) {
	v.Reset()
	poolAlibabaAscpPresalespackageConsignAPIRequest.Put(v)
}
