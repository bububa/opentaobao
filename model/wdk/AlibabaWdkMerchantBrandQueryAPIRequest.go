package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantBrandQueryAPIRequest 品牌查询接口 API请求
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
type AlibabaWdkMerchantBrandQueryAPIRequest struct {
	model.Params
	// 关键词，不填就查全部
	_keyword string
	// 可不填，默认0
	_offset int64
	// 可不填，默认2000
	_pageSize int64
}

// NewAlibabaWdkMerchantBrandQueryRequest 初始化AlibabaWdkMerchantBrandQueryAPIRequest对象
func NewAlibabaWdkMerchantBrandQueryRequest() *AlibabaWdkMerchantBrandQueryAPIRequest {
	return &AlibabaWdkMerchantBrandQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantBrandQueryAPIRequest) Reset() {
	r._keyword = ""
	r._offset = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.brand.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词，不填就查全部
func (r *AlibabaWdkMerchantBrandQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetOffset is Offset Setter
// 可不填，默认0
func (r *AlibabaWdkMerchantBrandQueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 可不填，默认2000
func (r *AlibabaWdkMerchantBrandQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaWdkMerchantBrandQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaWdkMerchantBrandQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantBrandQueryRequest()
	},
}

// GetAlibabaWdkMerchantBrandQueryRequest 从 sync.Pool 获取 AlibabaWdkMerchantBrandQueryAPIRequest
func GetAlibabaWdkMerchantBrandQueryAPIRequest() *AlibabaWdkMerchantBrandQueryAPIRequest {
	return poolAlibabaWdkMerchantBrandQueryAPIRequest.Get().(*AlibabaWdkMerchantBrandQueryAPIRequest)
}

// ReleaseAlibabaWdkMerchantBrandQueryAPIRequest 将 AlibabaWdkMerchantBrandQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantBrandQueryAPIRequest(v *AlibabaWdkMerchantBrandQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantBrandQueryAPIRequest.Put(v)
}
