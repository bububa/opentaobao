package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemDapeiTemplateQueryAPIRequest 搭配查询接口 API请求
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
type TmallItemDapeiTemplateQueryAPIRequest struct {
	model.Params
	// 搭配标题
	_title string
	// 页码
	_pageIndex int64
	// 分页大小
	_pageSize int64
}

// NewTmallItemDapeiTemplateQueryRequest 初始化TmallItemDapeiTemplateQueryAPIRequest对象
func NewTmallItemDapeiTemplateQueryRequest() *TmallItemDapeiTemplateQueryAPIRequest {
	return &TmallItemDapeiTemplateQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemDapeiTemplateQueryAPIRequest) Reset() {
	r._title = ""
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemDapeiTemplateQueryAPIRequest) GetApiMethodName() string {
	return "tmall.item.dapei.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemDapeiTemplateQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemDapeiTemplateQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 搭配标题
func (r *TmallItemDapeiTemplateQueryAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TmallItemDapeiTemplateQueryAPIRequest) GetTitle() string {
	return r._title
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *TmallItemDapeiTemplateQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallItemDapeiTemplateQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallItemDapeiTemplateQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallItemDapeiTemplateQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTmallItemDapeiTemplateQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemDapeiTemplateQueryRequest()
	},
}

// GetTmallItemDapeiTemplateQueryRequest 从 sync.Pool 获取 TmallItemDapeiTemplateQueryAPIRequest
func GetTmallItemDapeiTemplateQueryAPIRequest() *TmallItemDapeiTemplateQueryAPIRequest {
	return poolTmallItemDapeiTemplateQueryAPIRequest.Get().(*TmallItemDapeiTemplateQueryAPIRequest)
}

// ReleaseTmallItemDapeiTemplateQueryAPIRequest 将 TmallItemDapeiTemplateQueryAPIRequest 放入 sync.Pool
func ReleaseTmallItemDapeiTemplateQueryAPIRequest(v *TmallItemDapeiTemplateQueryAPIRequest) {
	v.Reset()
	poolTmallItemDapeiTemplateQueryAPIRequest.Put(v)
}
