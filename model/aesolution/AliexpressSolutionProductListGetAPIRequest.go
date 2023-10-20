package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductListGetAPIRequest Get product list API请求
// aliexpress.solution.product.list.get
//
// Get product list
type AliexpressSolutionProductListGetAPIRequest struct {
	model.Params
	// request parameters to query
	_aeopAEProductListQuery *ItemListQuery
}

// NewAliexpressSolutionProductListGetRequest 初始化AliexpressSolutionProductListGetAPIRequest对象
func NewAliexpressSolutionProductListGetRequest() *AliexpressSolutionProductListGetAPIRequest {
	return &AliexpressSolutionProductListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionProductListGetAPIRequest) Reset() {
	r._aeopAEProductListQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductListGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionProductListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAeopAEProductListQuery is AeopAEProductListQuery Setter
// request parameters to query
func (r *AliexpressSolutionProductListGetAPIRequest) SetAeopAEProductListQuery(_aeopAEProductListQuery *ItemListQuery) error {
	r._aeopAEProductListQuery = _aeopAEProductListQuery
	r.Set("aeop_a_e_product_list_query", _aeopAEProductListQuery)
	return nil
}

// GetAeopAEProductListQuery AeopAEProductListQuery Getter
func (r AliexpressSolutionProductListGetAPIRequest) GetAeopAEProductListQuery() *ItemListQuery {
	return r._aeopAEProductListQuery
}

var poolAliexpressSolutionProductListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionProductListGetRequest()
	},
}

// GetAliexpressSolutionProductListGetRequest 从 sync.Pool 获取 AliexpressSolutionProductListGetAPIRequest
func GetAliexpressSolutionProductListGetAPIRequest() *AliexpressSolutionProductListGetAPIRequest {
	return poolAliexpressSolutionProductListGetAPIRequest.Get().(*AliexpressSolutionProductListGetAPIRequest)
}

// ReleaseAliexpressSolutionProductListGetAPIRequest 将 AliexpressSolutionProductListGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionProductListGetAPIRequest(v *AliexpressSolutionProductListGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionProductListGetAPIRequest.Put(v)
}
