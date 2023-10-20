package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantCategoryQueryAPIRequest) Reset() {
	r._keyword = ""
	r._rootCategoryCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 搜索关键词，可不填就查全部
func (r *AlibabaWdkMerchantCategoryQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetRootCategoryCode is RootCategoryCode Setter
// 类目起点，可不填从根目录开始查
func (r *AlibabaWdkMerchantCategoryQueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
	r._rootCategoryCode = _rootCategoryCode
	r.Set("root_category_code", _rootCategoryCode)
	return nil
}

// GetRootCategoryCode RootCategoryCode Getter
func (r AlibabaWdkMerchantCategoryQueryAPIRequest) GetRootCategoryCode() string {
	return r._rootCategoryCode
}

var poolAlibabaWdkMerchantCategoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantCategoryQueryRequest()
	},
}

// GetAlibabaWdkMerchantCategoryQueryRequest 从 sync.Pool 获取 AlibabaWdkMerchantCategoryQueryAPIRequest
func GetAlibabaWdkMerchantCategoryQueryAPIRequest() *AlibabaWdkMerchantCategoryQueryAPIRequest {
	return poolAlibabaWdkMerchantCategoryQueryAPIRequest.Get().(*AlibabaWdkMerchantCategoryQueryAPIRequest)
}

// ReleaseAlibabaWdkMerchantCategoryQueryAPIRequest 将 AlibabaWdkMerchantCategoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantCategoryQueryAPIRequest(v *AlibabaWdkMerchantCategoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantCategoryQueryAPIRequest.Put(v)
}
