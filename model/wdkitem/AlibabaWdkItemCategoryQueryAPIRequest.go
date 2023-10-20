package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemcategoryqueryAPIRequest 类目查询接口 API请求
// alibaba.wdk.item.category.query
//
// 类目查询接口
type AlibabawdkitemcategoryqueryAPIRequest struct {
	model.Params
	// 查询关键词，不填查全部
	_keyword string
	// 从哪个类目开始查，不填从根类目开始查
	_rootCategoryCode string
}

// NewAlibabawdkitemcategoryqueryRequest 初始化AlibabawdkitemcategoryqueryAPIRequest对象
func NewAlibabawdkitemcategoryqueryRequest() *AlibabawdkitemcategoryqueryAPIRequest {
	return &AlibabawdkitemcategoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemcategoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemcategoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemcategoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 查询关键词，不填查全部
func (r *AlibabawdkitemcategoryqueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabawdkitemcategoryqueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetRootCategoryCode is RootCategoryCode Setter
// 从哪个类目开始查，不填从根类目开始查
func (r *AlibabawdkitemcategoryqueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
	r._rootCategoryCode = _rootCategoryCode
	r.Set("root_category_code", _rootCategoryCode)
	return nil
}

// GetRootCategoryCode RootCategoryCode Getter
func (r AlibabawdkitemcategoryqueryAPIRequest) GetRootCategoryCode() string {
	return r._rootCategoryCode
}
