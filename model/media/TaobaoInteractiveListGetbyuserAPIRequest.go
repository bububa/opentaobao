package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInteractiveListGetbyuserAPIRequest 用户获取视频互动列表 API请求
// taobao.interactive.list.getbyuser
//
// 根据用户来获取用户编辑的互动列表
type TaobaoInteractiveListGetbyuserAPIRequest struct {
	model.Params
	// 当前页
	_currentPage int64
	// 每页多少
	_pageSize int64
}

// NewTaobaoInteractiveListGetbyuserRequest 初始化TaobaoInteractiveListGetbyuserAPIRequest对象
func NewTaobaoInteractiveListGetbyuserRequest() *TaobaoInteractiveListGetbyuserAPIRequest {
	return &TaobaoInteractiveListGetbyuserAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInteractiveListGetbyuserAPIRequest) Reset() {
	r._currentPage = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetApiMethodName() string {
	return "taobao.interactive.list.getbyuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TaobaoInteractiveListGetbyuserAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页多少
func (r *TaobaoInteractiveListGetbyuserAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoInteractiveListGetbyuserAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInteractiveListGetbyuserRequest()
	},
}

// GetTaobaoInteractiveListGetbyuserRequest 从 sync.Pool 获取 TaobaoInteractiveListGetbyuserAPIRequest
func GetTaobaoInteractiveListGetbyuserAPIRequest() *TaobaoInteractiveListGetbyuserAPIRequest {
	return poolTaobaoInteractiveListGetbyuserAPIRequest.Get().(*TaobaoInteractiveListGetbyuserAPIRequest)
}

// ReleaseTaobaoInteractiveListGetbyuserAPIRequest 将 TaobaoInteractiveListGetbyuserAPIRequest 放入 sync.Pool
func ReleaseTaobaoInteractiveListGetbyuserAPIRequest(v *TaobaoInteractiveListGetbyuserAPIRequest) {
	v.Reset()
	poolTaobaoInteractiveListGetbyuserAPIRequest.Put(v)
}
