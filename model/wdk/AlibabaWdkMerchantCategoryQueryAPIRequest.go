package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantcategoryqueryAPIRequest 三江erp对接类目查询接口 API请求
// alibaba.wdk.merchant.category.query
//
// 三江erp对接类目查询接口
type AlibabawdkmerchantcategoryqueryAPIRequest struct {
	model.Params
	// 搜索关键词，可不填就查全部
	_keyword string
	// 类目起点，可不填从根目录开始查
	_rootCategoryCode string
}

// NewAlibabawdkmerchantcategoryqueryRequest 初始化AlibabawdkmerchantcategoryqueryAPIRequest对象
func NewAlibabawdkmerchantcategoryqueryRequest() *AlibabawdkmerchantcategoryqueryAPIRequest {
	return &AlibabawdkmerchantcategoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantcategoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantcategoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantcategoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 搜索关键词，可不填就查全部
func (r *AlibabawdkmerchantcategoryqueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabawdkmerchantcategoryqueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetRootCategoryCode is RootCategoryCode Setter
// 类目起点，可不填从根目录开始查
func (r *AlibabawdkmerchantcategoryqueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
	r._rootCategoryCode = _rootCategoryCode
	r.Set("root_category_code", _rootCategoryCode)
	return nil
}

// GetRootCategoryCode RootCategoryCode Getter
func (r AlibabawdkmerchantcategoryqueryAPIRequest) GetRootCategoryCode() string {
	return r._rootCategoryCode
}
