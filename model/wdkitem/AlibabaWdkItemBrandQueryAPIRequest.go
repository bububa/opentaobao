package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemBrandQueryAPIRequest 品牌信息查询 API请求
// alibaba.wdk.item.brand.query
//
// 品牌信息查询
type AlibabaWdkItemBrandQueryAPIRequest struct {
	model.Params
	// 查询关键词，不填则查询全部
	_keyword string
	// 起始位置
	_offset int64
	// 一页大小
	_pageSize int64
}

// NewAlibabaWdkItemBrandQueryRequest 初始化AlibabaWdkItemBrandQueryAPIRequest对象
func NewAlibabaWdkItemBrandQueryRequest() *AlibabaWdkItemBrandQueryAPIRequest {
	return &AlibabaWdkItemBrandQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemBrandQueryAPIRequest) Reset() {
	r._keyword = ""
	r._offset = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemBrandQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.brand.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemBrandQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemBrandQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 查询关键词，不填则查询全部
func (r *AlibabaWdkItemBrandQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaWdkItemBrandQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetOffset is Offset Setter
// 起始位置
func (r *AlibabaWdkItemBrandQueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r AlibabaWdkItemBrandQueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 一页大小
func (r *AlibabaWdkItemBrandQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaWdkItemBrandQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaWdkItemBrandQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemBrandQueryRequest()
	},
}

// GetAlibabaWdkItemBrandQueryRequest 从 sync.Pool 获取 AlibabaWdkItemBrandQueryAPIRequest
func GetAlibabaWdkItemBrandQueryAPIRequest() *AlibabaWdkItemBrandQueryAPIRequest {
	return poolAlibabaWdkItemBrandQueryAPIRequest.Get().(*AlibabaWdkItemBrandQueryAPIRequest)
}

// ReleaseAlibabaWdkItemBrandQueryAPIRequest 将 AlibabaWdkItemBrandQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemBrandQueryAPIRequest(v *AlibabaWdkItemBrandQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemBrandQueryAPIRequest.Put(v)
}
