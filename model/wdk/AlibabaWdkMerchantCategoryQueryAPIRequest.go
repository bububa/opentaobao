package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantCategoryQueryAPIRequest 三江erp对接类目查询接口 API请求
// alibaba.wdk.merchant.category.query
//
// 三江erp对接类目查询接口
type AlibabaWdkMerchantCategoryQueryAPIRequest struct {
	model.Params
	// 搜索关键词，可不填就查全部
	_keyword string
	// 类目起点，可不填从根目录开始查
	_rootCategoryCode string
}

// NewAlibabaWdkMerchantCategoryQueryRequest 初始化AlibabaWdkMerchantCategoryQueryAPIRequest对象
func NewAlibabaWdkMerchantCategoryQueryRequest() *AlibabaWdkMerchantCategoryQueryAPIRequest {
	return &AlibabaWdkMerchantCategoryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Keyword Setter
// 搜索关键词，可不填就查全部
func (r *AlibabaWdkMerchantCategoryQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// Get Keyword Getter
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// Set is RootCategoryCode Setter
// 类目起点，可不填从根目录开始查
func (r *AlibabaWdkMerchantCategoryQueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
	r._rootCategoryCode = _rootCategoryCode
	r.Set("root_category_code", _rootCategoryCode)
	return nil
}

// Get RootCategoryCode Getter
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetRootCategoryCode() string {
	return r._rootCategoryCode
}
