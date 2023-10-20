package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentradeusersgetAPIRequest 获取奇门用户列表 API请求
// taobao.qimen.trade.users.get
//
// 获取已开通奇门订单服务的用户列表
type TaobaoqimentradeusersgetAPIRequest struct {
	model.Params
	// 页数
	_pageSize int64
	// 每页的数量
	_pageIndex int64
}

// NewTaobaoqimentradeusersgetRequest 初始化TaobaoqimentradeusersgetAPIRequest对象
func NewTaobaoqimentradeusersgetRequest() *TaobaoqimentradeusersgetAPIRequest {
	return &TaobaoqimentradeusersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentradeusersgetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentradeusersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentradeusersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageSize is PageSize Setter
// 页数
func (r *TaobaoqimentradeusersgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoqimentradeusersgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 每页的数量
func (r *TaobaoqimentradeusersgetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoqimentradeusersgetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
