package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantCategoryQueryAPIRequest 查询商品的商家叶子类目 API请求
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
type AlibabaWdkItemMerchantCategoryQueryAPIRequest struct {
	model.Params
	// 请求
	_queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest
}

// NewAlibabaWdkItemMerchantCategoryQueryRequest 初始化AlibabaWdkItemMerchantCategoryQueryAPIRequest对象
func NewAlibabaWdkItemMerchantCategoryQueryRequest() *AlibabaWdkItemMerchantCategoryQueryAPIRequest {
	return &AlibabaWdkItemMerchantCategoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMerchantCategoryQueryAPIRequest) Reset() {
	r._queryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 请求
func (r *AlibabaWdkItemMerchantCategoryQueryAPIRequest) SetQueryRequest(_queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetQueryRequest() *WdkOpenSkuMerchantCatServiceQueryRequest {
	return r._queryRequest
}

var poolAlibabaWdkItemMerchantCategoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMerchantCategoryQueryRequest()
	},
}

// GetAlibabaWdkItemMerchantCategoryQueryRequest 从 sync.Pool 获取 AlibabaWdkItemMerchantCategoryQueryAPIRequest
func GetAlibabaWdkItemMerchantCategoryQueryAPIRequest() *AlibabaWdkItemMerchantCategoryQueryAPIRequest {
	return poolAlibabaWdkItemMerchantCategoryQueryAPIRequest.Get().(*AlibabaWdkItemMerchantCategoryQueryAPIRequest)
}

// ReleaseAlibabaWdkItemMerchantCategoryQueryAPIRequest 将 AlibabaWdkItemMerchantCategoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMerchantCategoryQueryAPIRequest(v *AlibabaWdkItemMerchantCategoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMerchantCategoryQueryAPIRequest.Put(v)
}
