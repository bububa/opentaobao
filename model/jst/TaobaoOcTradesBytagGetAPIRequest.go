package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaooctradesbytaggetAPIRequest 标签查询订单 API请求
// taobao.oc.trades.bytag.get
//
// 根据标签查询订单编号
type TaobaooctradesbytaggetAPIRequest struct {
	model.Params
	// 标签名称
	_tagName string
	// 当前页
	_page int64
	// 分页大小
	_pageSize int64
	// 标签类型，1官方，2自定义
	_tagType int64
}

// NewTaobaooctradesbytaggetRequest 初始化TaobaooctradesbytaggetAPIRequest对象
func NewTaobaooctradesbytaggetRequest() *TaobaooctradesbytaggetAPIRequest {
	return &TaobaooctradesbytaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaooctradesbytaggetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.trades.bytag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaooctradesbytaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaooctradesbytaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 标签名称
func (r *TaobaooctradesbytaggetAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r TaobaooctradesbytaggetAPIRequest) GetTagName() string {
	return r._tagName
}

// SetPage is Page Setter
// 当前页
func (r *TaobaooctradesbytaggetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaooctradesbytaggetAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaooctradesbytaggetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaooctradesbytaggetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetTagType is TagType Setter
// 标签类型，1官方，2自定义
func (r *TaobaooctradesbytaggetAPIRequest) SetTagType(_tagType int64) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// GetTagType TagType Getter
func (r TaobaooctradesbytaggetAPIRequest) GetTagType() int64 {
	return r._tagType
}
