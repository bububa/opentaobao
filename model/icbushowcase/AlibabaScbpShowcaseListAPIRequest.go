package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcaselistAPIRequest 橱窗查询 API请求
// alibaba.scbp.showcase.list
//
// 橱窗查询
type AlibabascbpshowcaselistAPIRequest struct {
	model.Params
	// 每页展示个数
	_perPageSize int64
	// 页码
	_toPage int64
}

// NewAlibabascbpshowcaselistRequest 初始化AlibabascbpshowcaselistAPIRequest对象
func NewAlibabascbpshowcaselistRequest() *AlibabascbpshowcaselistAPIRequest {
	return &AlibabascbpshowcaselistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcaselistAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcaselistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcaselistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPerPageSize is PerPageSize Setter
// 每页展示个数
func (r *AlibabascbpshowcaselistAPIRequest) SetPerPageSize(_perPageSize int64) error {
	r._perPageSize = _perPageSize
	r.Set("per_page_size", _perPageSize)
	return nil
}

// GetPerPageSize PerPageSize Getter
func (r AlibabascbpshowcaselistAPIRequest) GetPerPageSize() int64 {
	return r._perPageSize
}

// SetToPage is ToPage Setter
// 页码
func (r *AlibabascbpshowcaselistAPIRequest) SetToPage(_toPage int64) error {
	r._toPage = _toPage
	r.Set("to_page", _toPage)
	return nil
}

// GetToPage ToPage Getter
func (r AlibabascbpshowcaselistAPIRequest) GetToPage() int64 {
	return r._toPage
}
