package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCategoryQueryAPIRequest 类目查询接口 API请求
// alibaba.wdk.item.category.query
//
// 类目查询接口
type AlibabaWdkItemCategoryQueryAPIRequest struct {
	model.Params
	// 查询关键词，不填查全部
	_keyword string
	// 从哪个类目开始查，不填从根类目开始查
	_rootCategoryCode string
}

// NewAlibabaWdkItemCategoryQueryRequest 初始化AlibabaWdkItemCategoryQueryAPIRequest对象
func NewAlibabaWdkItemCategoryQueryRequest() *AlibabaWdkItemCategoryQueryAPIRequest {
	return &AlibabaWdkItemCategoryQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemCategoryQueryAPIRequest) Reset() {
	r._keyword = ""
	r._rootCategoryCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 查询关键词，不填查全部
func (r *AlibabaWdkItemCategoryQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetRootCategoryCode is RootCategoryCode Setter
// 从哪个类目开始查，不填从根类目开始查
func (r *AlibabaWdkItemCategoryQueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
	r._rootCategoryCode = _rootCategoryCode
	r.Set("root_category_code", _rootCategoryCode)
	return nil
}

// GetRootCategoryCode RootCategoryCode Getter
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetRootCategoryCode() string {
	return r._rootCategoryCode
}

var poolAlibabaWdkItemCategoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemCategoryQueryRequest()
	},
}

// GetAlibabaWdkItemCategoryQueryRequest 从 sync.Pool 获取 AlibabaWdkItemCategoryQueryAPIRequest
func GetAlibabaWdkItemCategoryQueryAPIRequest() *AlibabaWdkItemCategoryQueryAPIRequest {
	return poolAlibabaWdkItemCategoryQueryAPIRequest.Get().(*AlibabaWdkItemCategoryQueryAPIRequest)
}

// ReleaseAlibabaWdkItemCategoryQueryAPIRequest 将 AlibabaWdkItemCategoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemCategoryQueryAPIRequest(v *AlibabaWdkItemCategoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemCategoryQueryAPIRequest.Put(v)
}
