package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwuhomedecorationworkerlevellistAPIRequest 查询工人分层数据接口 API请求
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
type TmallfuwuhomedecorationworkerlevellistAPIRequest struct {
	model.Params
	// 当前页数
	_pageIndex int64
	// 每页大小，不传默认只查20条，最大不能超过500
	_pageSize int64
}

// NewTmallfuwuhomedecorationworkerlevellistRequest 初始化TmallfuwuhomedecorationworkerlevellistAPIRequest对象
func NewTmallfuwuhomedecorationworkerlevellistRequest() *TmallfuwuhomedecorationworkerlevellistAPIRequest {
	return &TmallfuwuhomedecorationworkerlevellistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfuwuhomedecorationworkerlevellistAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.homedecoration.workerlevel.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfuwuhomedecorationworkerlevellistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfuwuhomedecorationworkerlevellistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallfuwuhomedecorationworkerlevellistAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallfuwuhomedecorationworkerlevellistAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小，不传默认只查20条，最大不能超过500
func (r *TmallfuwuhomedecorationworkerlevellistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallfuwuhomedecorationworkerlevellistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
