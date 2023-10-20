package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSkuAttributeQueryAPIRequest Query the sku attribute information belonged to a specific category API请求
// aliexpress.solution.sku.attribute.query
//
// Query the sku attribute information belonged to a specific category, customized for oversea merchants.
type AliexpressSolutionSkuAttributeQueryAPIRequest struct {
	model.Params
	// input parameters
	_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest
}

// NewAliexpressSolutionSkuAttributeQueryRequest 初始化AliexpressSolutionSkuAttributeQueryAPIRequest对象
func NewAliexpressSolutionSkuAttributeQueryRequest() *AliexpressSolutionSkuAttributeQueryAPIRequest {
	return &AliexpressSolutionSkuAttributeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionSkuAttributeQueryAPIRequest) Reset() {
	r._querySkuAttributeInfoRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.sku.attribute.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuerySkuAttributeInfoRequest is QuerySkuAttributeInfoRequest Setter
// input parameters
func (r *AliexpressSolutionSkuAttributeQueryAPIRequest) SetQuerySkuAttributeInfoRequest(_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest) error {
	r._querySkuAttributeInfoRequest = _querySkuAttributeInfoRequest
	r.Set("query_sku_attribute_info_request", _querySkuAttributeInfoRequest)
	return nil
}

// GetQuerySkuAttributeInfoRequest QuerySkuAttributeInfoRequest Getter
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetQuerySkuAttributeInfoRequest() *SkuAttributeInfoQueryRequest {
	return r._querySkuAttributeInfoRequest
}

var poolAliexpressSolutionSkuAttributeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionSkuAttributeQueryRequest()
	},
}

// GetAliexpressSolutionSkuAttributeQueryRequest 从 sync.Pool 获取 AliexpressSolutionSkuAttributeQueryAPIRequest
func GetAliexpressSolutionSkuAttributeQueryAPIRequest() *AliexpressSolutionSkuAttributeQueryAPIRequest {
	return poolAliexpressSolutionSkuAttributeQueryAPIRequest.Get().(*AliexpressSolutionSkuAttributeQueryAPIRequest)
}

// ReleaseAliexpressSolutionSkuAttributeQueryAPIRequest 将 AliexpressSolutionSkuAttributeQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionSkuAttributeQueryAPIRequest(v *AliexpressSolutionSkuAttributeQueryAPIRequest) {
	v.Reset()
	poolAliexpressSolutionSkuAttributeQueryAPIRequest.Put(v)
}
