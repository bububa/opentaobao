package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantCategoryQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
