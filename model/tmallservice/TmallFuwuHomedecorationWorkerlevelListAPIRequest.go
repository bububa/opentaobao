package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationWorkerlevelListAPIRequest 查询工人分层数据接口 API请求
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
type TmallFuwuHomedecorationWorkerlevelListAPIRequest struct {
	model.Params
	// 当前页数
	_pageIndex int64
	// 每页大小，不传默认只查20条，最大不能超过500
	_pageSize int64
}

// NewTmallFuwuHomedecorationWorkerlevelListRequest 初始化TmallFuwuHomedecorationWorkerlevelListAPIRequest对象
func NewTmallFuwuHomedecorationWorkerlevelListRequest() *TmallFuwuHomedecorationWorkerlevelListAPIRequest {
	return &TmallFuwuHomedecorationWorkerlevelListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFuwuHomedecorationWorkerlevelListAPIRequest) Reset() {
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuHomedecorationWorkerlevelListAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.homedecoration.workerlevel.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuHomedecorationWorkerlevelListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFuwuHomedecorationWorkerlevelListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallFuwuHomedecorationWorkerlevelListAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallFuwuHomedecorationWorkerlevelListAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小，不传默认只查20条，最大不能超过500
func (r *TmallFuwuHomedecorationWorkerlevelListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallFuwuHomedecorationWorkerlevelListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTmallFuwuHomedecorationWorkerlevelListAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFuwuHomedecorationWorkerlevelListRequest()
	},
}

// GetTmallFuwuHomedecorationWorkerlevelListRequest 从 sync.Pool 获取 TmallFuwuHomedecorationWorkerlevelListAPIRequest
func GetTmallFuwuHomedecorationWorkerlevelListAPIRequest() *TmallFuwuHomedecorationWorkerlevelListAPIRequest {
	return poolTmallFuwuHomedecorationWorkerlevelListAPIRequest.Get().(*TmallFuwuHomedecorationWorkerlevelListAPIRequest)
}

// ReleaseTmallFuwuHomedecorationWorkerlevelListAPIRequest 将 TmallFuwuHomedecorationWorkerlevelListAPIRequest 放入 sync.Pool
func ReleaseTmallFuwuHomedecorationWorkerlevelListAPIRequest(v *TmallFuwuHomedecorationWorkerlevelListAPIRequest) {
	v.Reset()
	poolTmallFuwuHomedecorationWorkerlevelListAPIRequest.Put(v)
}
