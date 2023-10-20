package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaointeractivelistgetbyuserAPIRequest 用户获取视频互动列表 API请求
// taobao.interactive.list.getbyuser
//
// 根据用户来获取用户编辑的互动列表
type TaobaointeractivelistgetbyuserAPIRequest struct {
	model.Params
	// 当前页
	_currentPage int64
	// 每页多少
	_pageSize int64
}

// NewTaobaointeractivelistgetbyuserRequest 初始化TaobaointeractivelistgetbyuserAPIRequest对象
func NewTaobaointeractivelistgetbyuserRequest() *TaobaointeractivelistgetbyuserAPIRequest {
	return &TaobaointeractivelistgetbyuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaointeractivelistgetbyuserAPIRequest) GetApiMethodName() string {
	return "taobao.interactive.list.getbyuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaointeractivelistgetbyuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaointeractivelistgetbyuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TaobaointeractivelistgetbyuserAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaointeractivelistgetbyuserAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页多少
func (r *TaobaointeractivelistgetbyuserAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaointeractivelistgetbyuserAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
