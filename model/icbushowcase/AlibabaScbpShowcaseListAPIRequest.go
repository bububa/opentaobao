package icbushowcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseListAPIRequest 橱窗查询 API请求
// alibaba.scbp.showcase.list
//
// 橱窗查询
type AlibabaScbpShowcaseListAPIRequest struct {
	model.Params
	// 每页展示个数
	_perPageSize int64
	// 页码
	_toPage int64
}

// NewAlibabaScbpShowcaseListRequest 初始化AlibabaScbpShowcaseListAPIRequest对象
func NewAlibabaScbpShowcaseListRequest() *AlibabaScbpShowcaseListAPIRequest {
	return &AlibabaScbpShowcaseListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseListAPIRequest) Reset() {
	r._perPageSize = 0
	r._toPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPerPageSize is PerPageSize Setter
// 每页展示个数
func (r *AlibabaScbpShowcaseListAPIRequest) SetPerPageSize(_perPageSize int64) error {
	r._perPageSize = _perPageSize
	r.Set("per_page_size", _perPageSize)
	return nil
}

// GetPerPageSize PerPageSize Getter
func (r AlibabaScbpShowcaseListAPIRequest) GetPerPageSize() int64 {
	return r._perPageSize
}

// SetToPage is ToPage Setter
// 页码
func (r *AlibabaScbpShowcaseListAPIRequest) SetToPage(_toPage int64) error {
	r._toPage = _toPage
	r.Set("to_page", _toPage)
	return nil
}

// GetToPage ToPage Getter
func (r AlibabaScbpShowcaseListAPIRequest) GetToPage() int64 {
	return r._toPage
}

var poolAlibabaScbpShowcaseListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseListRequest()
	},
}

// GetAlibabaScbpShowcaseListRequest 从 sync.Pool 获取 AlibabaScbpShowcaseListAPIRequest
func GetAlibabaScbpShowcaseListAPIRequest() *AlibabaScbpShowcaseListAPIRequest {
	return poolAlibabaScbpShowcaseListAPIRequest.Get().(*AlibabaScbpShowcaseListAPIRequest)
}

// ReleaseAlibabaScbpShowcaseListAPIRequest 将 AlibabaScbpShowcaseListAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseListAPIRequest(v *AlibabaScbpShowcaseListAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseListAPIRequest.Put(v)
}
