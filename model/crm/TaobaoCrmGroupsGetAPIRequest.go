package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupsgetAPIRequest 查询卖家的分组 API请求
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
type TaobaocrmgroupsgetAPIRequest struct {
	model.Params
	// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
	_pageSize int64
	// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
	_currentPage int64
}

// NewTaobaocrmgroupsgetRequest 初始化TaobaocrmgroupsgetAPIRequest对象
func NewTaobaocrmgroupsgetRequest() *TaobaocrmgroupsgetAPIRequest {
	return &TaobaocrmgroupsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgroupsgetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.groups.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgroupsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgroupsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageSize is PageSize Setter
// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
func (r *TaobaocrmgroupsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaocrmgroupsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
func (r *TaobaocrmgroupsgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaocrmgroupsgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
